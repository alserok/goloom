package pages

import (
	"html/template"
)

const (
	statePage = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>🪼 Goloom</title>
    <style>
        body {
			font-family: 'Montserrat', sans-serif;
            background-color: #1a1a1a; /* Dark background */
            margin: 0;
            padding: 0;
            display: flex;
			flex-direction: column;
            align-items: center;
            min-height: 100vh;
            color: #e0e0e0; /* Light text color */
        }

        .navbar {
            background-color: #2d2d2d; /* Dark navbar background */
            padding: 10px 20px;
            display: flex;
            justify-content: space-between;
            align-items: center;
            border-radius: 25px; /* Rounded borders */
            margin: 20px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
			gap: 20px;
			width: 70vw;
        }

        /* Navbar brand/logo */
        .navbar-brand {
            font-size: 24px;
            font-weight: bold;
            color: #ffffff; /* White text */
            text-decoration: none;
        }

        /* Navbar links */
.navbar-links {
            display: flex;
            gap: 20px; /* Space between links */
        }

        .navbar-links a {
            color: #e0e0e0; /* Light text color */
            text-decoration: none;
            padding: 5px 18px;
            border-radius: 20px; /* Rounded borders for links */
            transition: background-color 0.3s ease, color 0.3s ease;
        }

        .navbar-links a:hover {
            background-color: #007bff; /* Blue background on hover */
            color: white; /* White text on hover */
        }

        /* Responsive design */
        @media (max-width: 600px) {
            .navbar {
                flex-direction: column;
                align-items: flex-start;
                padding: 10px;
            }

            .navbar-links {
                flex-direction: column;
                gap: 10px;
                width: 100%%;
            }

            .navbar-links a {
                width: 100%%;
                text-align: center;
            }
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
            width: 70vw;
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
            border-radius: 50%;
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

<nav class="navbar">
    <div href="#" class="navbar-brand">🪼 Goloom</div>
    <div class="navbar-links">
        <a href="/web/config/dir/"><i class="fas fa-home"></i>Config</a>
    </div>
</nav>

<div class="container">
    <button class="refresh-button" onclick="window.location.reload()">↻</button>
    <h1>🫧 Services</h1>
	<h3 style="color:gray">Updated at: {{ .time }}</h3>
    <table>
        <thead>
            <tr>
                <th>Service</th>
                <th>Status</th>
            </tr>
        </thead>
        <tbody>
			{{range .states}}
				{{ if eq .Status 200 }}
					<tr>
						<td>{{ .Service }}</td>
						<td><span class="status up">OK</span></td>
					</tr>
				{{else}}
					<tr>
						<td>{{ .Service }}</td>
						<td><span class="status down">DOWN</span></td>
					</tr>
        		{{end}}
    		{{end}}
        </tbody>
    </table>
</div>

 <script>
        setInterval(async function() {
             window.location.reload()
        }, 10_000);
    </script>
</body>
</html>`
)

func newStatePage() *template.Template {
	page, err := template.New("state").Parse(statePage)
	if err != nil {
		panic("failed to generate page: " + err.Error())
	}

	return page
}
