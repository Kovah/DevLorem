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
 *
 */

// Split the URL into segments
$url = array_values(array_filter(explode('/', $_SERVER["REQUEST_URI"])));

// Process the API if requested
if (isset($url[0]) && $url[0] === 'api') {

    $content = '';

    $quotes = getRandomQuotes();
    $split_content = $quotes['content'];
    $source = $quotes['source'];

    if (isset($url[1]) && preg_match("/[0-9]/", $url[1])) {
        $split_content = fillOrTrimQuotes($split_content, $url[1]);
    }

    if ((isset($url[2]) && $url[2] === "json") || (isset($url[3]) && $url[3] === "json")) {
        $json_content = array();

        // Process the content for JSON output
        foreach ($split_content as $paragraph) {
            if (!empty($paragraph)) {
                // Check if the p tags should be visible
                if (isset($url[2]) && $url[2] === "p") {
                    array_push($json_content, $paragraph);
                } else {
                    array_push($json_content, preg_replace("/(\<(\/)?p\>)/", "", $paragraph));
                }
            }
        }

        print_r(json_encode($json_content));

    } else {

        // Process the content
        foreach ($split_content as $paragraph) {
            if (!empty($paragraph)) {
                // Check if the p tags should be visible
                if (isset($url[2]) && $url[2] === "p") {
                    $content .= $paragraph . ' ';
                } else {
                    $content .= preg_replace("/(\<(\/)?p\>)/", "", $paragraph);
                }
            }
        }

        print_r($content);
    }

    exit;

}