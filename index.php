<?php
define("DEVLOREM", true);

// Load the functions
require_once('functions.php');

// Load the API engine
require_once('api.php');

// Define needed variables
$menu_link = "/p";
$menu_text = "Show Paragraphs";
$headline = "";
$content = "";

$quotes = getRandomQuotes();
$split_content = $quotes['content'];
$headline = $quotes['source'];

// Process the content
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
