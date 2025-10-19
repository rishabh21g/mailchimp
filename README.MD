# 📧 Mailchimp Clone - Concurrent Email Sender

A high-performance email sending system built with Go that demonstrates the producer-consumer pattern using goroutines and channels. Send personalized HTML emails to multiple recipients concurrently using local SMTP testing with Mailpit.

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Active-success?style=for-the-badge)

## 📸 Screenshots

### Architecture & Data Flow
<p align="center">
  <img src="https://github.com/rishabh21g/mailchimp/blob/main/images/MailchimpArch.png" alt="Architecture Diagram" width="800">
</p>

### Mailpit Dashboard - All Emails Delivered
<p align="center">
  <img src="https://github.com/rishabh21g/mailchimp/blob/main/images/MailchimpGo.png" alt="Mailpit Inbox" width="800">
</p>
*Successfully sent 20 personalized emails to all recipients*

### Email Preview - Personalized Content
<p align="center">
  <img src="https://github.com/rishabh21g/mailchimp/blob/main/images/Inbox_temp.png" alt="Email Preview" width="800">
</p>
*Dynamic HTML template with recipient-specific data (Name and Email)*

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

### 5. Expected Output

```bash
Worker 1: Sending email to: John.Smith@example.com
Worker 2: Sending email to: Sarah.Johnson@email.com
Worker 3: Sending email to: Michael.Brown@test.com
Worker 4: Sending email to: Emily.Davis@sample.org
Worker 5: Sending email to: David.Wilson@demo.net
...
```

## 📁 Project Structure

```
mailchimp/
├── images/               # Screenshots and diagrams
│   ├── architecture.png
│   ├── mailpit-inbox.png
│   └── email-preview.png
├── main.go              # Entry point and worker pool initialization
├── producer.go          # CSV reader and channel producer
├── consumer.go          # Email worker implementation
├── template.go          # Template parsing and execution
├── types.go             # Data structures
├── email.tmpl           # HTML email template
├── email.csv            # Recipient data
└── README.md            # Project documentation
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

# 🚀 Future Improvements

## 📊 Database Integration
- [ ] Store contacts in PostgreSQL/MySQL instead of CSV
- [ ] Contact metadata (phone, company, location, tags)
- [ ] Track email delivery status and logs
- [ ] Soft delete and archive functionality

## 🎨 Frontend Development
- [ ] Build web dashboard (React/Next.js + TailwindCSS)
- [ ] Contact management interface with search/filter
- [ ] Campaign creation wizard
- [ ] Real-time sending progress bar
- [ ] Email preview before sending
- [ ] Analytics dashboard (open rates, clicks, bounces)

## 📅 Email Scheduling
- [ ] Schedule campaigns for future delivery
- [ ] Recurring campaigns (daily, weekly, monthly)
- [ ] Timezone-aware sending
- [ ] Drip campaigns with delays
- [ ] Pause/resume functionality

## 🎨 Template System
- [ ] Drag-and-drop visual template builder
- [ ] Pre-built template library (newsletter, promo, transactional)
- [ ] Template versioning and storage
- [ ] Mobile-responsive preview
- [ ] Dynamic personalization tokens

## 🔧 Error Handling
- [ ] Retry logic with exponential backoff
- [ ] Track and categorize failures (hard bounce, soft bounce)
- [ ] Email validation (syntax, DNS/MX records)
- [ ] Duplicate email prevention
- [ ] Real-time error alerts (Slack/Discord)

## 👥 Contact Grouping
- [ ] Create and manage contact groups
- [ ] Dynamic segments based on criteria
- [ ] Tag system for contacts
- [ ] Bulk operations (import, export, tag)
- [ ] Engagement-based segmentation

## 🔐 Authentication & Security
- [ ] User authentication with JWT
- [ ] Role-based access control (Admin, Editor, Viewer)
- [ ] API key management
- [ ] Rate limiting and CSRF protection

## 📈 Analytics
- [ ] Track open rates (pixel tracking)
- [ ] Click-through rate (link tracking)
- [ ] Campaign performance charts
- [ ] Export reports to PDF/CSV

## 🔌 Integrations
- [ ] SMTP providers (SendGrid, AWS SES, Mailgun)
- [ ] Webhook notifications
- [ ] RESTful API with documentation
- [ ] CRM integrations

## 📱 Additional Features
- [ ] Unsubscribe management & preference center
- [ ] Double opt-in for subscribers
- [ ] A/B testing for subject lines
- [ ] Multi-language support

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