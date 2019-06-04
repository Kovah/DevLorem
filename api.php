<?php
if (!defined('DEVLOREM')) {
    exit('No direct script access allowed');
}

/*
 * DevLorem API
 *
 * Required URL structure:
 * domain.com/api/[int][/p][/json]
 *
 * [int] = number of paragraphs you want
 * [/p] = select if the <p> tags should be included
 * [/json] = output the data in JSON format
 *
 * Example:
 * domain.com/api/4             will output 4 paragraphs
 * domain.com/api/4/p           will output 4 paragraphs with <p> tags
 * domain.com/api/4/json        will output 4 json-encoded paragraphs
 * domain.com/api/4/p/json      will output 4 json-encoded paragraphs with <p> tags
 *
 */

// Split the URL into segments
$url = array_values(array_filter(explode('/', $_SERVER['REQUEST_URI'])));

// Process the API if requested
if (isset($url[0]) && $url[0] === 'api') {

    $content = '';

    $quotes = getRandomQuotes();
    $split_content = $quotes['content'];
    $source = $quotes['source'];

    $outputJson = (isset($url[2]) && $url[2] === 'json') || (isset($url[3]) && $url[3] === 'json');
    $outputParagraphs = isset($url[2]) && $url[2] === 'p';

    if (isset($url[1]) && preg_match('[0-9].', $url[1])) {
        $split_content = fillOrTrimQuotes($split_content, $url[1]);
    }

    if ($outputJson) {
        $json_content = array();

        // Process the content for JSON output
        foreach ($split_content as $paragraph) {
            if (!empty($paragraph)) {
                // Check if the p tags should be visible
                if ($outputParagraphs) {
                    $json_content[] = $paragraph;
                } else {
                    $json_content[] = preg_replace("/(\<(\/)?p\>)/", '', $paragraph);
                }
            }
        }

        header('Content-Type: application/json');
        echo json_encode($json_content);

    } else {

        // Process the content
        foreach ($split_content as $paragraph) {
            if (!empty($paragraph)) {
                // Check if the p tags should be visible
                if ($outputParagraphs) {
                    $content .= $paragraph . ' ';
                } else {
                    $content .= preg_replace("/(\<(\/)?p\>)/", '', $paragraph);
                }
            }
        }

        header('Content-Type: text/plain');
        echo $content;
    }

    exit;

}
