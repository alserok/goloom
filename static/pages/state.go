package pages

import (
	"context"
	"fmt"
	"github.com/alserok/goloom/internal/service/models"
	"net/http"
	"strings"
	"time"
)

const (
	statePage = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Service Status</title>
    <style>
        /* Dark theme base styles */
        body {
            font-family: 'Arial', sans-serif;
            background-color: #1a1a1a; /* Dark background */
            margin: 0;
            padding: 0;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            color: #e0e0e0; /* Light text color */
        }

        .container {
            background-color: #2d2d2d; /* Dark container background */
            border-radius: 12px;
            box-shadow: 0 8px 16px rgba(0, 0, 0, 0.3);
            padding: 24px;
            width: 90%%;
            max-width: 800px;
            position: relative; /* For positioning the refresh button */
        }

        h1 {
            text-align: center;
            color: #ffffff; /* White heading */
            margin-bottom: 24px;
            font-size: 28px;
            font-weight: 600;
        }

        h3 {
            text-align: center;
            color: #ffffff; /* White heading */
            margin-bottom: 24px;
            font-size: 16px;
            font-weight: 400;
        }

        table {
            width: 100%%;
            border-collapse: collapse;
            margin-top: 16px;
        }

        th, td {
            padding: 16px;
            text-align: left;
            border-bottom: 1px solid #444444; /* Darker border */
        }

        th {
            background-color: #007bff; /* Blue header */
            color: white;
            font-weight: 600;
            text-transform: uppercase;
            letter-spacing: 1px;
        }

        tr {
            transition: background-color 0.3s ease;
        }

        tr:hover {
            background-color: #3a3a3a; /* Darker hover effect */
        }

        .status {
            padding: 6px 12px;
            border-radius: 20px;
            display: inline-block;
            text-align: center;
            font-size: 14px;
			width: 4vw;
        }

        .status.up {
            background-color: #28a745; /* Green for "up" status */
            color: white;
        }

        .status.down {
            background-color: #dc3545; /* Red for "down" status */
            color: white;
        }

        /* Refresh button */
        .refresh-button {
            position: absolute;
            top: 20px;
            right: 20px;
            background-color: gray; /* Blue to match the header */
            color: white;
            border: none;
            border-radius: 50%%;
            width: 40px;
            height: 40px;
            font-size: 18px;
            cursor: pointer;
            display: flex;
            align-items: center;
            justify-content: center;
            transition: background-color 0.3s ease;
        }

        .refresh-button:hover {
           opacity: 0.9;
        }

        /* Animations */
        @keyframes fadeIn {
            from {
                opacity: 0;
                transform: translateY(-20px);
            }
            to {
                opacity: 1;
                transform: translateY(0);
            }
        }

        /* Responsive design */
        @media (max-width: 600px) {
            .container {
                width: 95%%;
                padding: 16px;
            }

            h1 {
                font-size: 24px;
            }

            th, td {
                padding: 12px;
            }

            .refresh-button {
                top: 10px;
                right: 10px;
                width: 36px;
                height: 36px;
                font-size: 16px;
            }
        }
    </style>
</head>
<body>

<div class="container">
    <button class="refresh-button" onclick="window.location.reload()">↻</button>
    <h1>Goloom</h1>
	<h3 style="color:gray">Updated at: %s</h3>
    <table>
        <thead>
            <tr>
                <th>Service</th>
                <th>Status</th>
            </tr>
        </thead>
        <tbody>
			%s
        </tbody>
    </table>
</div>

</body>
</html>`
)

func NewStatePage(ctx context.Context, states []models.ServiceState) (string, error) {
	var sb strings.Builder

	for _, state := range states {
		switch state.Status {
		case http.StatusOK:
			sb.WriteString(fmt.Sprintf(`
			<tr>
				<td>%s</td>
				<td><span class="status up">OK</span></td>
			</tr>
		`, state.Service))
		default:
			sb.WriteString(fmt.Sprintf(`
			<tr>
				<td>%s</td>
				<td><span class="status down">DOWN</span></td>
			</tr>
		`, state.Service))
		}
	}

	page := fmt.Sprintf(statePage, time.Now().Format("2006-01-02 15:04:05"), sb.String())

	return page, nil
}
