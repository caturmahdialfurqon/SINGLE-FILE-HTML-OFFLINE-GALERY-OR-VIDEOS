package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Video struct {
	Title string
	Source string
}

func main() {
	var directory, titledisplay, fileName string

	fmt.Print("Enter Directory Path Here => ")
	fmt.Scanln(&directory)

	fmt.Print("Enter Title Bar for Your HTML file => ")
	fmt.Scanln(&titledisplay)

	// Check if the directory exists
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		fmt.Printf("Error: Directory '%s' does not exist.\n", directory)
		return
	}

	// Open the directory for reading
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error: Could not read directory '%s'.\n", directory)
		return
	}

	// Initialize an empty list to store video data
	var videoList []Video

	// Loop through each file in the directory
	for _, file := range files {
		// Skip hidden files (starting with a dot)
		if file.IsDir() || strings.HasPrefix(file.Name(), ".") || !strings.HasSuffix(strings.ToLower(file.Name()), ".mp4") {
			continue
		}

		// Extract video title from filename (can be improved for better handling)
		title := strings.TrimSuffix(file.Name(), ".mp4")

		// Create a video item with title and source path
		video := Video{
			Title: title,
			Source: filepath.Join(directory, file.Name()),
		}

		// Check if the video is already in the list (case-insensitive)
		found := false
		for _, item := range videoList {
			if strings.ToLower(item.Title) == strings.ToLower(title) {
				found = true
				break
			}
		}

		// Add the video item to the list if not a duplicate
		if !found {
			videoList = append(videoList, video)
		}
	}

	// Build the complete HTML structure
	html := `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>` + titledisplay + `</title>
    <style>
        body {
            background-color: #000033;
            color: rgb(255, 255, 255, 0.95);
            font-family: Arial, sans-serif;
            height: 100vh; /* Set body height to 100% of viewport height */
            display: flex; /* Enable flexbox for the body */
            flex-direction: column; /* Arrange elements vertically */
            align-items: center; /* Center items horizontally */
        }

        .container {
            width: auto;
            margin: 0; /* Remove default margin for better layout */
            display: flex; /* Use flexbox for column layout */
            flex-wrap: wrap; /* Allow items to wrap to multiple lines */
            flex: 1; /* Allow container to expand to fill available space */
        }

        .video-list {
            /* Remove width and padding restrictions */
            width: auto;
            padding: 0; /* Remove default padding for better spacing */
            list-style-type: none;
            display: flex; /* Use flexbox for column layout */
            flex-wrap: wrap; /* Allow items to wrap to multiple lines */
            flex: 1; /* Allow video list to expand to fill available space */
        }

        .video-list li {
            /* Remove margin and padding restrictions */
            margin: 0;
            padding: 0; /* Remove default padding for better spacing */
            border: 1px solid #00BFFF; /* Add a subtle border */
            border-radius: 5px; /* Add rounded corners */
            width: calc(100% / 3); /* Define width for each video item (adjust for desired columns) */
            box-sizing: border-box; /* Include padding and border in width calculation */
            flex: 0 0 auto; /* Prevent items from stretching beyond their natural size */
            aspect-ratio: 4/3; /* Set aspect ratio for videos (16:9 is common for most videos) */
            overflow: hidden; /* Hide any content that overflows the aspect ratio */
        }

        .video-title {
            font-size: 18px;
            margin-bottom: 10px;
        }

        video {
            width: 100%;
            height: auto;
        }
        
    </style>
</head>
<body>
  <div class="container">
    <ul class="video-list">`

	// Add video list items using the same loop from before
	for _, video := range videoList {
		html += `<li>
            <div class="video-title">` + video.Title + `</div>
            <video controls>
              <source src="` + video.Source + `" type="video/mp4">
              Your browser does not support the video tag.
            </video>
          </li>`
	}

	html += `</ul>
  </div>
</body>
</html>`

	// Define the filename for the generated HTML file
	fmt.Print("Enter your file name *without .html* => ")
	fmt.Scanln(&fileName)
	outputPath := filepath.Join(directory, fileName+".html")

	// Open the file for writing (overwrite existing content)
	err = ioutil.WriteFile(outputPath, []byte(html), 0644)
	if err != nil {
		fmt.Printf("Error: Could not open file '%s' for writing.\n", outputPath)
		return
	}

	// Success message
	fmt.Printf("Generated HTML file saved to: %s\n", outputPath)
}

