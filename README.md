# go-feedback-classification-api
Backend service for classifying customer feedback using rule-based sentiment analysis. Designed for future AI integration.
# 🤖 AI Feedback Classification API

Backend service that classifies customer feedback into Positive or Negative categories.  
Designed with extensibility in mind for future AI/ML integration.

## 🚀 Tech Stack
- Golang (Gin Framework)
- RESTful API
- Clean Architecture

## 🧱 Architecture
Handler → Service → Repository

## ✨ Features
- Submit customer feedback
- Rule-based sentiment classification
- Retrieve feedback by sentiment
- Simple and extensible design for future AI integration

## 🧠 Classification Logic
- Keyword-based sentiment scoring
- Easy to replace with ML model later

## ▶️ Run Application

### Run locally
```bash
go run ./cmd/server
## ▶️ Run Application

Health Check
GET http://localhost:8080/health
