# Groupie Tracker

## 📌 Project Overview
Groupie Tracker is a web application that fetches and displays data about music artists, their concerts, and related information from an API. Users can browse through artist details and concert locations in a structured and interactive manner.

## 📁 Project Structure
The project follows a modular structure to ensure clean and maintainable code:

```
GROUPIE-TRACKER/
│── cmd/
│   └── main.go              # Entry point of the application
│
├── handler/                 # Handles HTTP requests
│   ├── Detail_Func.go       # Handles artist details page
│   ├── Fetch_Data.go        # Fetches data from API
│   ├── Groupie_Func.go      # Handles main page request
│   ├── Style_Func.go        # Manages styles for the website
│
├── helpers/
│   └── renderTemplates.go   # Utility for rendering templates
│
├── static/
│   └── style.css            # Styling for the website
│
├── template/                # HTML templates
│   ├── detailsCard.html     # Template for artist details
│   ├── index.html           # Homepage template
│   ├── statusPage.html      # Error/status page template
│
├── tools/
│   └── Tools.go             # Contains data structures and utility functions
│
├── go.mod                   # Go module file
└── README.md                # Project documentation
```

## 🔧 Features
- Fetch and display artist information dynamically.
- View artist details, including name, image, members, and concerts.
- Handle 404 and error pages gracefully.
- Clean and structured project organization.

## 🚀 Installation & Setup
### Prerequisites
- Install [Go](https://go.dev/)

### Steps
1. Clone this repository:
   ```sh
   git clone https://github.com/yourusername/groupie-tracker.git
   cd groupie-tracker
   ```
2. Initialize Go modules:
   ```sh
   go mod tidy
   ```
3. Run the project:
   ```sh
   go run cmd/main.go
   ```
4. Open your browser and visit:
   ```
   http://localhost:8080
   ```

## 📡 API Source
This project fetches data from the [Groupie Tracker API](https://groupietrackers.herokuapp.com/api/).

## 🛠 Technologies Used
- **Go**: Backend development
- **HTML & CSS**: Frontend
- **net/http**: Server and routing
- **encoding/json**: Handling API responses
- **text/template**: HTML rendering

## 📝 License
This project is open-source and available under the [MIT License](LICENSE).

