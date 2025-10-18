# 📧 Mailchimp Clone - Concurrent Email Sender

A high-performance email sending system built with Go that demonstrates the producer-consumer pattern using goroutines and channels. Send personalized HTML emails to multiple recipients concurrently using local SMTP testing with Mailpit.

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Active-success?style=for-the-badge)

## 🎯 Features

- ✨ **Concurrent Email Processing** - Multiple worker goroutines for parallel email sending
- 📊 **Producer-Consumer Pattern** - Efficient channel-based communication
- 🎨 **HTML Email Templates** - Beautiful, responsive email designs using Go templates
- 📝 **CSV Data Import** - Load recipient data from CSV files
- 🔧 **Local SMTP Testing** - Integration with Mailpit for development
- ⚡ **High Performance** - Rate-limited concurrent workers (5 workers by default)
- 🎭 **Personalized Content** - Dynamic template rendering with recipient data

## 🏗️ Architecture

```
┌─────────────┐
│  email.csv  │
└──────┬──────┘
       │
       ▼
┌─────────────────┐        ┌──────────────┐
│   Producer      │───────▶│   Channel    │
│ LoadRecipients  │        │ (Unbuffered) │
└─────────────────┘        └──────┬───────┘
                                  │
                    ┌─────────────┼─────────────┐
                    ▼             ▼             ▼
              ┌──────────┐  ┌──────────┐  ┌──────────┐
              │ Worker 1 │  │ Worker 2 │  │ Worker N │
              └────┬─────┘  └────┬─────┘  └────┬─────┘
                   │             │             │
                   └─────────────┼─────────────┘
                                 ▼
                        ┌─────────────────┐
                        │  SMTP Server    │
                        │  (Mailpit)      │
                        │  localhost:1025 │
                        └─────────────────┘
```

## 📋 Prerequisites

- Go 1.21 or higher
- Docker (for Mailpit SMTP server)

## 🚀 Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/rishabh21g/mailchimp.git
cd mailchimp
```

### 2. Start Mailpit SMTP Server

```bash
docker run -d \
  --restart unless-stopped \
  --name=mailpit \
  -p 8025:8025 \
  -p 1025:1025 \
  axllent/mailpit
```

Access Mailpit web interface at: **http://localhost:8025**

### 3. Prepare Your Data

Create or update `email.csv` with your recipient data:

```csv
Name,Email
John Smith,John.Smith@example.com
Sarah Johnson,Sarah.Johnson@email.com
```

### 4. Run the Application

```bash
go run .
```

## 📁 Project Structure

```
mailchimp/
├── main.go           # Entry point and worker pool initialization
├── producer.go       # CSV reader and channel producer
├── consumer.go       # Email worker implementation
├── template.go       # Template parsing and execution
├── types.go          # Data structures
├── email.tmpl        # HTML email template
├── email.csv         # Recipient data
└── README.md         # Project documentation
```

## 🎨 Email Template

The project uses Go's `html/template` package for rendering personalized emails. The template includes:

- Responsive HTML design
- Personalized greeting with recipient name
- Dynamic email address insertion
- Professional styling with CSS
- Call-to-action button
- Footer with unsubscribe link

### Template Variables

```go
type Recipient struct {
    Name  string
    Email string
}
```

## ⚙️ Configuration

### Adjust Worker Count

Modify `main.go`:

```go
workers := 5 // Change this value
```

### SMTP Settings

Modify `consumer.go`:

```go
smtpHost := "localhost"
smtpPort := "1025"
```

### Rate Limiting

Adjust delay between emails in `consumer.go`:

```go
time.Sleep(time.Microsecond * 50) // Adjust as needed
```

## 📊 Performance

- **Workers**: 5 concurrent goroutines
- **Rate**: ~50 microseconds delay between sends
- **Scalability**: Easily adjustable based on SMTP server limits

## 🔍 Monitoring

View sent emails in real-time:

1. Open your browser
2. Navigate to `http://localhost:8025`
3. Watch emails arrive in the Mailpit inbox

## 🛠️ Troubleshooting

### Emails not sending?

1. Ensure Mailpit is running:
   ```bash
   docker ps | grep mailpit
   ```

2. Check SMTP connection:
   ```bash
   telnet localhost 1025
   ```

3. Review application logs for errors

### Template parsing errors?

- Verify `email.tmpl` exists in the project root
- Check template syntax for proper Go template tags
- Ensure proper MIME headers in template

## 📝 To-Do / Roadmap

- [ ] Add email authentication support
- [ ] Implement retry logic for failed sends
- [ ] Add progress bar/status indicator
- [ ] Support for attachments
- [ ] Email queue management
- [ ] Database integration
- [ ] Configuration file support
- [ ] Metrics and analytics
- [ ] Email scheduling

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 👨‍💻 Author

**Rishabh Gupta**

- GitHub: [@rishabh21g](https://github.com/rishabh21g)
- Email: rishabhiitm@zohomail.in

## 🙏 Acknowledgments

- [Mailpit](https://github.com/axllent/mailpit) - SMTP testing tool
- Go community for excellent concurrency patterns
- All contributors and users of this project

## 📚 Learn More

- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Producer-Consumer Pattern](https://en.wikipedia.org/wiki/Producer%E2%80%93consumer_problem)
- [SMTP Protocol](https://datatracker.ietf.org/doc/html/rfc5321)

---

⭐ **Star this repository if you find it helpful!**

```
Made with ❤️ and Go
```