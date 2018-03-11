package main

import ("log"
        "net/smtp"
)

type EmailUser struct {
  UserName string
  Password string
  EmailServer string
  Port int
}

func main()  {

  emailUser := &EmailUser{"someemailaddress", "somesecretpwd", "smtp.gmail.com", 587}

  auth := smtp.PlainAuth(
    "",
    emailUser.UserName,
    emailUser.Password,
    emailUser.EmailServer,
  )

  err := smtp.SendMail(
    "smtp.gmail.com:587",
    auth,
    "recipient@email.com",
    []string{"recipient@email.com"},
    []byte("This is the email body"),
  )

  if err != nil {
    log.Fatal(err)
  }

}
