<?php
define("DEVLOREM", true);

// Define needed variables
$menu_link = "/p";
$menu_text = "Show Paragraphs";
$headline = "";
$content = "";

// Select a random lorem file
$dir = __DIR__ . "/lorem/";
$files = glob($dir . "*.*");
$file = array_rand($files);

// Set output variables
$content = file_get_contents($files[$file]);
$headline = str_replace($dir, "", $files[$file]);
$headline = strtoupper(str_replace(".txt", "", $headline));

// Process the content
$split_content = explode(PHP_EOL, $content);
$content = "";
foreach ($split_content as $paragraph) {
    if (!empty($paragraph)) {
        // Check if the p tags should be visible
        if ($_SERVER["REQUEST_URI"] === "/p") {
            $content .= "<p>" . htmlspecialchars($paragraph) . "</p>";

            // Change the menulink to plain
            $menu_link = "/";
            $menu_text = "Hide Paragraphs";
        } else {
            $content .= $paragraph;
        }
    }
}

// Load the template
require_once('template.php');
