package pages

import (
	"html/template"
)

var (
	configPage = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>🪼 Goloom: Edit file</title>
    <style>
        /* Dark theme base styles */
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
			width: 70vw;
            display: flex;
            justify-content: space-between;
            align-items: center;
            border-radius: 25px; /* Rounded borders */
            margin: 20px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
			gap: 20px;
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
            width: 70vw;
            max-width: 800px;
            animation: fadeIn 0.5s ease-in-out;
            position: relative; /* For positioning the save button */
        }

        h1 {
            text-align: center;
            color: #ffffff; /* White heading */
            margin-bottom: 24px;
            font-size: 28px;
            font-weight: 600;
        }

        textarea {
            height: 400px;
			width: 95%;
            background-color: #3a3a3a; /* Darker textarea background */
            color: #e0e0e0; /* Light text color */
            border-radius: 8px;
            padding: 12px;
            font-family: 'Courier New', monospace; /* Monospace font for code */
            font-size: 14px;
            outline: none; /* Remove default outline */
			resize: none;
        }

        textarea:focus {
            border-color: #007bff; /* Blue border on focus */
        }

        .save-button {
            display: none; /* Hidden by default */
            position: absolute;
            bottom: -60px; /* Position below the container */
            right: 0;
            background-color: #007bff; /* Blue button */
            color: white;
            border: none;
            border-radius: 20px;
            padding: 10px 20px;
            font-size: 16px;
            cursor: pointer;
            transition: background-color 0.3s ease;
        }

        .save-button.visible {
            display: block; /* Show when content changes */
        }

        .save-button:hover {
            background-color: #0056b3; /* Darker blue on hover */
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

            textarea {
                height: 300px;
            }
        }

	.delete-button {
		display: none; /* Hidden by default */
		position: absolute;
		bottom: -60px; /* Position below the container */
		left: 0px; /* Adjust position to avoid overlap with save button */
		background-color: #ff4d4d; /* Red button */
		color: white;
		border: none;
		border-radius: 20px;
		padding: 10px 20px;
		font-size: 16px;
		cursor: pointer;
		transition: background-color 0.3s ease;
	}
	
	.delete-button.visible {
		display: block; /* Show when content is loaded */
	}
	
	.delete-button:hover {
		background-color: #cc0000; /* Darker red on hover */
	}

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
    <h1>{{ .filename }}</h1>
    <textarea id="editor" spellcheck="false"></textarea>
	<button id="delete-button" class="delete-button">Delete</button>
    <button id="save-button" class="save-button">Save</button>
</div>

<script>
    const editor = document.getElementById('editor');
    const saveButton = document.getElementById('save-button');

	const initialContent = ('{{ .content }}'.trim() === '{}' || '{{ .content }}'.trim() === '') ? '' : '{{ .content }}';
	editor.value = initialContent;

    let isContentChanged = false;

    editor.addEventListener('input', () => {
        if (!isContentChanged) {
            isContentChanged = true;
            saveButton.classList.add('visible');
        }
    });

    function goBack() {
		const fallbackUrl = "/web/config/dir/"; // Fallback URL
		const currentPath = window.location.pathname;
	
		const modifiedPath = currentPath.replace("/web/config/file", "/web/config/dir");
	
		const normalizedPath = modifiedPath.endsWith("/") ? modifiedPath.slice(0, -1) : modifiedPath;
	
		const pathParts = normalizedPath.split("/");
		const parentPath = pathParts.slice(0, -1).join("/") || fallbackUrl;
	
		window.location.href = parentPath;
    }

    saveButton.addEventListener('click', async () => {
        const updatedContent = editor.value;

        try {
            const response = await fetch('/config/update/{{ .path }}', {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ config: updatedContent }),
            });

            if (response.ok) {
                isContentChanged = false;
                saveButton.classList.remove('visible');
            }
        } catch (error) {
            console.error('Error:', error);
        }
    });

	const deleteButton = document.getElementById('delete-button');
	
	deleteButton.classList.add('visible');
	
	deleteButton.addEventListener('click', async () => {
			try {
				const response = await fetch('/config/delete/{{ .path }}', {
					method: 'DELETE',
				});
	
				if (response.ok) {
					alert('File deleted successfully!');
					window.location.href = "/web/config/dir";
				} else {
					alert('Failed to delete the file.');
				}
			} catch (error) {
				console.error('Error:', error);
				alert('An error occurred while deleting the file.');
			}
	});
</script>

</body>
</html>`
)

func newConfigPage() *template.Template {
	page, err := template.New("config").Parse(configPage)
	if err != nil {
		panic("failed to generate page: " + err.Error())
	}

	return page
}
