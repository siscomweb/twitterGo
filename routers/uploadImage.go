package routers

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"strings"
	"twitterGo/basedatos"
	"twitterGo/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type readSeeker struct {
	io.Reader
}

func (rs *readSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func UploadImage(ctx context.Context, uploadType string, request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400
	IDUsuario := claim.ID.Hex()

	var filename string
	var usuario models.Usuario

	fmt.Println("UploadImage: bucket")
	//bucket := aws.String(ctx.Value(models.Key("bucketName")).(string))
	bucket := aws.String("bucketName")
	switch uploadType {
	case "A":
		filename = "avatars/" + IDUsuario + ".jpg"
		usuario.Avatar = filename
	case "B":
		filename = "banners/" + IDUsuario + ".jpg"
		usuario.Banner = filename
	}

	fmt.Println("UploadImage: mediaType")
	mediaType, params, err := mime.ParseMediaType(request.Headers["Content-Type"])
	if err != nil {
		r.Status = 500
		r.Message = err.Error()
		return r
	}

	fmt.Println("UploadImage: multipart")
	if strings.HasPrefix(mediaType, "multipart/") {
		body, err := base64.StdEncoding.DecodeString(request.Body)
		if err != nil {
			r.Status = 500
			r.Message = err.Error()
			return r
		}

		fmt.Println("UploadImage: boundary")
		mr := multipart.NewReader(bytes.NewReader(body), params["boundary"])
		p, err := mr.NextPart()
		if err != nil && err != io.EOF {
			r.Status = 500
			r.Message = err.Error()
			return r
		}

		if err != io.EOF {
			if p.FileName() != "" {
				buff := bytes.NewBuffer(nil)
				if _, err := io.Copy(buff, p); err != nil {
					r.Status = 500
					r.Message = err.Error()
					return r
				}

				fmt.Println("UploadImage: NewSession AWS")
				sess, err := session.NewSession(&aws.Config{
					Region: aws.String("us-east-1")})
				if err != nil {
					r.Status = 500
					r.Message = err.Error()
					return r
				}

				fmt.Println("UploadImage: s3manager")
				uploader := s3manager.NewUploader(sess)
				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: bucket,
					Key:    aws.String(filename),
					Body:   &readSeeker{buff},
				})
				if err != nil {
					r.Status = 500
					r.Message = err.Error()
					return r
				}
			}
		}

		fmt.Println("UploadImage: ModificoRegistro")
		status, err := basedatos.ModificoRegistro(usuario, IDUsuario)
		if err != nil || !status {
			r.Status = 500
			r.Message = "Error al modificar registro del usuario. >" + err.Error()
			return r
		}
	} else {
		r.Message = "Debe enviar una immagen con el 'Content-Type' de tipo 'multipart/' en el Header"
		r.Status = 400
		return r
	}

	r.Status = 200
	r.Message = "Image upoload OK!"
	return r
}
