# Website Uptime Monitor (CLI + Web)

**A lightweight and versatile website uptime monitoring tool written in Go, supporting both command-line and web dashboard modes.**

Monitor your websites in real-time, log their status, and optionally visualize uptime history via a simple web interface.

---

## **Features**
- ✅ Monitor websites via CLI or file input  
- ✅ Concurrent HTTP checks for multiple URLs  
- ✅ Logs results to SQLite database  
- ✅ Web dashboard to view recent status history  
- ✅ Configurable check intervals  
- ✅ Single-binary deployment, lightweight and portable

---

## **Table of Contents**
- [Installation](#installation)  
- [Usage](#usage)  
  - [CLI Mode](#cli-mode)  
  - [Web Server Mode](#web-server-mode)  
- [Configuration](#configuration)  
- [Database](#database)  
- [Extending the Tool](#extending-the-tool)  
- [License](#license)

---

## **Installation**

### **Prerequisites**
- Go 1.25+ installed  
- SQLite3 (optional, for persistent logging)

### **Build**
Clone the repository and build the binary:

```bash
git clone https://github.com/yourusername/uptime-monitor.git
cd uptime-monitor
go build -o uptime-monitor
```

You will get a single executable: `uptime-monitor`.

---

## **Usage**

### **CLI Mode**
Monitor URLs directly from the command line or a file:

```bash
# Monitor specific URLs every 60 seconds (default)
./uptime-monitor https://google.com https://example.com

# Set custom interval (in seconds)
./uptime-monitor -interval 30 https://example.com

# Load URLs from a file
./uptime-monitor --file urls.txt
```

**Output:**  
```
[UP] https://google.com
[DOWN] https://example.com : 500 Internal Server Error
```

---

### **Web Server Mode**
Start a web dashboard to view recent uptime logs:

```bash
./uptime-monitor --web
```

- Dashboard runs on **http://localhost:8000**  
- Displays last 50 checks, including URL, status (UP/DOWN), timestamp, and error info if applicable.

---

## **Configuration**
- `-interval <seconds>` – Time between checks (default: 60s)  
- `--file <path>` – Path to file with URLs (one per line)  
- `--web` – Start as web server instead of CLI  

**Example combining file input with interval:**
```bash
./uptime-monitor -interval 30 --file urls.txt
```

---

## **Database**
- Uses **SQLite** to log website status.  
- Database location: `./db/uptime.db`  
- Table structure:  
```sql
CREATE TABLE IF NOT EXISTS status_logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    url TEXT,
    up BOOLEAN,
    info TEXT,
    checked_at DATETIME
);
```
- Each check inserts a new row with URL, status, optional error info, and timestamp.

---

## **Extending the Tool**
- Add **Slack or Email notifications** for downtime alerts.  
- Implement **uptime percentage calculation** per website.  
- Build **frontend dashboard enhancements** with charts or filtering.  
- Schedule the tool with `systemd` or `cron` for continuous monitoring.  
- Deploy as a lightweight Docker container for cloud hosting.

---

## **Deployment**
1. Build the binary:  
```bash
go build -o uptime-monitor
```
2. Initialize the SQLite database if needed.  
3. Run in CLI or Web mode.  
4. Optionally, run as a service or container for continuous uptime monitoring.

---

## **License**
This project is licensed under the MIT License – see the [LICENSE](LICENSE) file for details.

---
