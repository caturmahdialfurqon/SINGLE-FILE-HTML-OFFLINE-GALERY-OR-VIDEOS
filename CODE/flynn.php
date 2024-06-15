<?php

// Define the directory containing the video files
//$directory = "/Users/furqonflynn/Movies/FETHBIT/FLORIDAGAS";
$directory = readline("Enter Directory Path Here => ");
$titledisplay = readline("Enter Tittle Bar for Your HTML file =>");

// Open the directory for reading
if (!is_dir($directory)) {
  die("Error: Directory '$directory' does not exist.");
}

$handle = opendir($directory);

// Initialize an empty list to store video data
$videoList = [];

// Loop through each file in the directory
while (($file = readdir($handle)) !== false) {

  // Skip non-video files and current/parent directory entries
  if (in_array($file, ['.', '..']) || !preg_match('/\.mp4$/i', $file)) {
    continue;
  }

  // Extract video title from filename (can be improved for better handling)
  $title = str_replace(".mp4", "", $file);

  // Create a video item with title and source path
  $video = [
    "title" => $title,
    "source" => "$directory/$file"
  ];

  // Check if the video is already in the list (case-insensitive)
  $found = false;
  foreach ($videoList as $item) {
    if (strtolower($item["title"]) === strtolower($title)) {
      $found = true;
      break;
    }
  }

  // Add the video item to the list if not a duplicate
  if (!$found) {
    $videoList[] = $video;
  }
}

// Close the directory handle
closedir($handle);

// Build the complete HTML structure
$html = "<!DOCTYPE html>
<html lang=\"en\">
<head>
  <meta charset=\"UTF-8\">
  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">
  <title>$titledisplay</title>
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
  <div class=\"container\">
    <ul class=\"video-list\">";

// Add video list items using the same loop from before
foreach ($videoList as $video) {
  $html .= "<li>
                <div class=\"video-title\">{$video["title"]}</div>
                <video controls>
                  <source src=\"{$video["source"]}\" type=\"video/mp4\">
                  Your browser does not support the video tag.
                </video>
              </li>";
}

$html .= "</ul>
  </div>
</body>
</html>";

// Define the filename for the generated HTML file
$fileNamex = readline("Enter your file name *without .html* => ");
$fileName = "$fileNamex.html";

// Get the current script's directory (optional)
//$scriptDir = dirname(__FILE__);

// Combine script directory and filename
$filePath = "$directory/$fileName";

$fileHandle = fopen($filePath, "w") or die("Error: Could not open file '$filePath' for writing.");

// Open the file for writing (overwrite existing content)
//$fileHandle = fopen($fileName, "w") or die("Error: Could not open file '$fileName' for writing.");

// Write the complete HTML content to the file
fwrite($fileHandle, $html);

// Close the file handle
fclose($fileHandle);

// Success message
echo "Generated HTML file saved to: $filePath";

?>

