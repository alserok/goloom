package pages

import (
	"html/template"
)

const (
	dirPage = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>🪼 Goloom</title>
  <style>
    /* General Styles */
    body {
      font-family: 'Arial', sans-serif;
      background-color: #1a1a1a; /* Dark background */
      margin: 0;
      padding: 0;
      display: flex;
      flex-direction: column;
      align-items: center;
      min-height: 100vh;
      color: #e0e0e0; /* Light text color */
    }

    /* Navbar Styles */
    .navbar {
      background-color: #2d2d2d; /* Dark navbar background */
      padding: 10px 20px;
      width: 70vw;
      max-width: 800px;
      display: flex;
      justify-content: space-between;
      align-items: center;
      border-radius: 25px; /* Rounded borders */
      margin: 20px 0;
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
      gap: 20px;
    }

    .navbar-brand {
      font-size: 24px;
      font-weight: bold;
      color: #ffffff; /* White text */
      text-decoration: none;
    }

    .navbar-links {
      display: flex;
      gap: 20px; /* Space between links */
    }

    .navbar-links a {
      color: #e0e0e0; /* Light text color */
      text-decoration: none;
      padding: 8px 16px;
      border-radius: 20px; /* Rounded borders for links */
      transition: background-color 0.3s ease, color 0.3s ease;
    }

    .navbar-links a:hover {
      background-color: #007bff; /* Blue background on hover */
      color: white; /* White text on hover */
    }

    /* Container Styles */
    .container {
      background-color: #2d2d2d;
      padding: 20px;
      border-radius: 12px;
      width: 70vw;
      max-width: 800px;
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
    }

    h1 {
      color: #ffffff;
      font-size: 16px;
      margin-bottom: 20px;
      color: gray;
    }

    /* Item Styles */
    .item {
      background-color: #3d3d3d;
      padding: 15px;
      margin-bottom: 10px;
      border-radius: 8px;
      display: flex;
      justify-content: space-between;
      align-items: center;
      transition: background-color 0.3s ease, transform 0.2s ease;
      cursor: pointer;
    }

    .item:hover {
      background-color: #4d4d4d;
      transform: translateY(-2px);
    }

    .item a {
      color: #e0e0e0;
      text-decoration: none;
      font-weight: bold;
      flex-grow: 1;
    }

    .item a:hover {
      color: #ffffff;
      text-decoration: underline;
    }

    .item .type {
      color: #888;
      font-style: italic;
      font-size: 14px;
	  font-weight: bold;
      margin-left: 25px;
    }

    .item .size {
      color: white;
      font-size: 14px;
	  font-weight: bold;
      margin-left: 15px;
    }

    /* Back Button Styles */
    .back-button {
      background-color: #007bff;
      color: white;
      border: none;
      padding: 10px 20px;
      border-radius: 5px;
      cursor: pointer;
      font-size: 16px;
      margin-bottom: 20px;
      transition: background-color 0.3s ease;
    }

    .back-button:hover {
      opacity: 0.9;
    }

    /* Responsive Design */
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

      .container {
        padding: 15px;
      }
    }
  </style>
</head>
<body>
  <nav class="navbar">
    <div href="#" class="navbar-brand">🪼 Goloom</div>
    <div class="navbar-links">
      <a href="/web/state"><i class="fas fa-home"></i>State</a>
    </div>
  </nav>

  <div class="container">
    <button class="back-button" onclick="goBack()">Back</button>
    <h1>{{ .path }}</h1>
    <div id="content"></div>
  </div>

  <script>
    const data = {{ .dir }};

    function renderMap(map) {
      const content = document.getElementById("content");

      for (const [key, value] of Object.entries(map)) {
        const div = document.createElement("div");
        div.className = "item";

        const link = value.is_dir ? '/web/config/dir/' + {{ .path }} + "/" + value.name : '/web/config/file/' + {{ .path }} + "/" + value.name;

        const anchor = document.createElement("a");
        anchor.href = link;
        anchor.textContent = value.is_dir ? value.name : value.name.split('.').slice(0, -1);

        const size = document.createElement("span");
        size.className = "size";
        size.textContent = value.size;

        const type = document.createElement("span");
        type.className = "type";
        type.textContent = value.is_dir ? '📁' : (value.name.includes('.') ? value.name.split('.').pop() : '') + ' 📃';

        div.appendChild(anchor);
        div.appendChild(size);
        div.appendChild(type);

        content.appendChild(div);
      }
    }

    // Function to handle "Go Back" button
    function goBack() {
      const fallbackUrl = "/"; // Fallback URL if no parent directory exists
      const currentPath = window.location.pathname;
      const parentPath = currentPath.split("/").slice(0, -1).join("/") || fallbackUrl;
      window.location.href = parentPath;
    }

    // Render the map
    renderMap(data.content);
  </script>
</body>
</html>
`
)

func newDirPage() *template.Template {
	page, err := template.New("dir").Parse(dirPage)
	if err != nil {
		panic("failed to generate page: " + err.Error())
	}

	return page
}
