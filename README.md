# Devflow-Console 💻

A dashboard console that provides real-time repository insights, PR analytics, and test reports through visual depiction of certain kinds of charts. Built with Golang and GitHub GraphQL API, it optimizes data fetching for an efficient developer experience.

## ✨ Features  
- 📊 **Repository Metrics** – Displays total commits, stars, and other key insights.  
- 🔍 **Template based Insights** – Fetches relevant data from GitHub API and shows it accordingly per template. Templates are organized as :
  1. *General Insights* – Shows general insights like total commits, stars, forks, etc.
  2. *Developer Insights* – Shows insights related to developers like total PRs, commits, etc.
  3. *Qa Engineer Insights* – Shows insights related to QA engineers like total test cases, test cases passed, etc.
  4. *Project Manager Insights* – Shows insights related to project managers like total issues, issues closed, etc.




- 🛠 **Versatile Visuals** – Displays certain aspects of the metrics through different kinds of charts
- 🚀 **Optimized API Calls** – Uses GraphQL for efficient data retrieval.  
- 📦 **Modular & Scalable Backend** – Structured Go backend for maintainability.  

## 🏗️ Tech Stack  
- **Frontend:** Next.js, ShadCN UI, Recharts
- **Backend:** Golang , Gin
- **Data Fetching:** GitHub GraphQL API   

## ⚡ Getting Started  

### Prerequisites  
- Go installed  
- GitHub Personal Access Token (for API requests)  

### Installation  
1. Clone the repository:  
   ```bash
   git clone https://github.com/ayyush08/devflow-console.git
   cd devflow-console
   ```
2. Create a `.env` file in the **client** directory and add the following environment variables:  
   ```bash
    NEXT_PUBLIC_BACKEND_URL=your-backend-url
    ```

3. Create a `.env` file in the **server** directory and add the following environment variables:
    ```bash
    PORT=port to run the server, make sure it is not 3000, as 3000 is for client
    GITHUB_ACCESS_TOKEN=your-github-access-token to fetch data 
    FRONTEND_URL=your-frontend-url
    ```

4. Run the backend server:  
   ```bash
   cd server
   go run main.go
   ```
5. Run the frontend server:  
   ```bash
   cd client
    npm install
    npm run dev
    ```

6. Visit `http://localhost:3000` to view the application.

## Contributing and Feedback
Any kind of contributions are welcome. Please feel free to fork the repository and make a pull request. For feedback, suggestions, or queries, feel free to open an issue.
