<?php
if (!defined('DEVLOREM')) {
    exit('No direct script access allowed');
}

function getRandomQuotes() {
    // Select a random lorem file
    $dir = __DIR__ . "/lorem/";
    $files = glob($dir . "*.*");
    $file = array_rand($files);

    // Get the file contents and format the source
    $content = file_get_contents($files[$file]);
    $source = strtoupper(str_replace(".txt", "", str_replace($dir, "", $files[$file])));

    $split_content = explode(PHP_EOL, $content);
    shuffle($split_content);

    return array(
        'content' => $split_content,
        'source' => $source,
    );
}

function fillOrTrimQuotes($split_content, $count) {
    while (count($split_content) < $count) {
        $quotes = getRandomQuotes();
        $split_content = array_merge($split_content, $quotes['content']);
    }

    // Trim the array if there are too much quotes
    if (count($split_content) > $count) {
        $split_content = array_slice($split_content, 0, $count);
    }

    return $split_content;
}