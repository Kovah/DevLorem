<?php
if (!defined('DEVLOREM')) {
    exit('No direct script access allowed');
}
?>
<!doctype html>

<html lang="en">
<head>
    <meta charset="utf-8">
    <title>DevLorem</title>
    <link rel="stylesheet" href="styles.css?v=1.0">
</head>
<body>

<div class="container">

    <div class="menu">
        <ul>
            <li><a href="<?php echo $menu_link; ?>"><?php echo $menu_text; ?></a></li>
            <li>(Click a paragraph three times to select the text)</li>
        </ul>
    </div>

    <div class="content">
        <h1>
            <?php echo $headline; ?>
            <button onclick="window.location.href=window.location.href">
                <span>Reload</span>
                <img src="data:image/png;base64,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" alt="Reload">
            </button>
        </h1>
        <?php echo (empty($content) ? 'Sorry, an error occurred!' : $content); ?>
    </div>

    <div class="footer">
        A Project by <a href="https://kovah.de">Kovah.de</a> | Also available <a href="https://devlorem.kovah.de">online</a> |
        <a href="https://github.com/Kovah/DevLorem">Contribute</a>
    </div>

</div>

</body>
</html>