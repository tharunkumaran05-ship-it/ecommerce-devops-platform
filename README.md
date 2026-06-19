```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestsTotal.Inc()

	html := `
<!DOCTYPE html>
<html>
<head>
    <title>CommerceOps Platform</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <style>
        *{
            margin:0;
            padding:0;
            box-sizing:border-box;
            font-family:Arial,sans-serif;
        }

        body{
            background:linear-gradient(135deg,#0f172a,#1e293b);
            color:white;
            min-height:100vh;
            padding:40px;
        }

        .container{
            max-width:1200px;
            margin:auto;
        }

        .hero{
            text-align:center;
            margin-bottom:40px;
        }

        .hero h1{
            font-size:3rem;
            margin-bottom:10px;
        }

        .hero p{
            color:#cbd5e1;
            font-size:1.2rem;
        }

        .grid{
            display:grid;
            grid-template-columns:repeat(auto-fit,minmax(250px,1fr));
            gap:20px;
            margin-top:30px;
        }

        .card{
            background:#1e293b;
            border-radius:16px;
            padding:25px;
            text-align:center;
            box-shadow:0 4px 15px rgba(0,0,0,0.3);
        }

        .card h3{
            margin-bottom:10px;
        }

        .status{
            color:#22c55e;
            font-weight:bold;
        }

        .buttons{
            margin-top:35px;
            text-align:center;
        }

        .btn{
            display:inline-block;
            background:#3b82f6;
            color:white;
            text-decoration:none;
            padding:12px 24px;
            border-radius:8px;
            margin:10px;
            transition:0.3s;
        }

        .btn:hover{
            background:#2563eb;
        }

        .section{
            margin-top:40px;
            background:#1e293b;
            padding:25px;
            border-radius:16px;
        }

        .section h2{
            margin-bottom:15px;
        }

        ul{
            padding-left:20px;
            line-height:2;
        }

        footer{
            text-align:center;
            margin-top:40px;
            color:#94a3b8;
        }
    </style>
</head>

<body>

<div class="container">

    <div class="hero">
        <h1>🚀 CommerceOps Platform</h1>
        <p>End-to-End DevOps Platform using Go, Kubernetes, GitOps & Monitoring</p>
    </div>

    <div class="grid">

        <div class="card">
            <h3>🐳 Docker</h3>
            <p class="status">Containerized</p>
        </div>

        <div class="card">
            <h3>🏗 Terraform</h3>
            <p class="status">Kubernetes Resources</p>
        </div>

        <div class="card">
            <h3>☸ Kubernetes</h3>
            <p class="status">Running</p>
        </div>

        <div class="card">
            <h3>📦 Helm</h3>
            <p class="status">Deployed</p>
        </div>

        <div class="card">
            <h3>⚙ GitHub Actions</h3>
            <p class="status">CI/CD Active</p>
        </div>

        <div class="card">
            <h3>🔄 Argo CD</h3>
            <p class="status">GitOps Enabled</p>
        </div>

        <div class="card">
            <h3>🔥 Prometheus</h3>
            <p class="status">Metrics Scraping</p>
        </div>

        <div class="card">
            <h3>📊 Grafana</h3>
            <p class="status">Dashboards Active</p>
        </div>

    </div>

    <div class="buttons">
        <a class="btn" href="/health">Health Check</a>
        <a class="btn" href="/metrics">Metrics</a>
    </div>

    <div class="section">
        <h2>Deployment Workflow</h2>

        <p>
        Developer → GitHub → GitHub Actions →
        Docker Hub → Argo CD →
        Kubernetes → Prometheus → Grafana
        </p>

        <br>

        <p>
        Terraform → Kubernetes Infrastructure
        </p>
    </div>

    <div class="section">
        <h2>Project Features</h2>

        <ul>
            <li>Go-based Web Application</li>
            <li>Docker Multi-Stage Containerization</li>
            <li>Kubernetes Deployment & Service Management</li>
            <li>Helm Chart Packaging</li>
            <li>GitHub Actions CI/CD Automation</li>
            <li>GitOps Deployment using Argo CD</li>
            <li>Infrastructure as Code using Terraform</li>
            <li>Prometheus Metrics Collection</li>
            <li>Grafana Monitoring Dashboards</li>
            <li>Custom HTTP Request Metrics (http_requests_total)</li>
        </ul>
    </div>

    <div class="section">
        <h2>Tech Stack</h2>

        <ul>
            <li>Go (Golang)</li>
            <li>Docker</li>
            <li>Kubernetes</li>
            <li>Helm</li>
            <li>Terraform</li>
            <li>GitHub Actions</li>
            <li>Argo CD</li>
            <li>Prometheus</li>
            <li>Grafana</li>
        </ul>
    </div>

    <footer>
        <p>Built by Richard Benjamin</p>
        <p>Aspiring DevOps Engineer | Kubernetes | Monitoring | Automation</p>
    </footer>

</div>

</body>
</html>
`

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}
```
