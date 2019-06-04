<?php
if (!defined('DEVLOREM')) {
    exit('No direct script access allowed');
}
?>
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>DevLorem for <?php echo $headline; ?></title>
    <link rel="stylesheet" href="/styles.css?v=2.0">
</head>
<body>

<div class="container">

    <div class="menu">
        <ul>
            <li><a href="<?php echo $menu_link; ?>"><?php echo $menu_text; ?></a></li>
            <li>(Click a paragraph to copy the text)</li>
        </ul>
    </div>

    <div class="content">
        <h1>
            <?php echo $headline; ?>
            <button onclick="window.location.href=window.location.href" title="Reload the page for new texts">
                <span>Reload</span>
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                    <path fill="#1998a6"
                        d="M500.33 0h-47.41a12 12 0 0 0-12 12.57l4 82.76A247.42 247.42 0 0 0 256 8C119.34 8 7.9 119.53 8 256.19 8.1 393.07 119.1 504 256 504a247.1 247.1 0 0 0 166.18-63.91 12 12 0 0 0 .48-17.43l-34-34a12 12 0 0 0-16.38-.55A176 176 0 1 1 402.1 157.8l-101.53-4.87a12 12 0 0 0-12.57 12v47.41a12 12 0 0 0 12 12h200.33a12 12 0 0 0 12-12V12a12 12 0 0 0-12-12z"></path>
                </svg>
            </button>
        </h1>
        <?php echo(empty($content) ? 'Sorry, an error occurred!' : $content); ?>
    </div>

    <div class="footer">
        A Project by <a href="https://kovah.de">Kovah.de</a> | Also available <a
            href="https://devlorem.kovah.de">online</a>
        and
        <a href="#api" id="api-toggle">via API</a> | <a href="https://corporatelorem.kovah.de">Lorem for your
            clients</a> |
        <a href="https://github.com/Kovah/DevLorem">Contribute</a>
    </div>

    <div id="api" class="footer" style="display: none;">
        <p>
            Required URL structure:<br/>
            <code><?php echo($_SERVER["HTTP_HOST"] ?? 'yourdomain.com'); ?>
                /api/[int][/p][/json]</code>
        </p>
        <ul>
            <li>[int] = optional, number of paragraphs you want</li>
            <li>[/p] = optional, select if the <code>&lt;p&gt;</code> tags should be included</li>
            <li>[/json] = optional, output the data in JSON format</li>
        </ul>
    </div>

</div>

<script src="/clipboard.min.js"></script>
<script>
    document.getElementById('api-toggle').addEventListener('click', function (e) {
        e.preventDefault();
        var $apiWrapper = document.getElementById('api');
        if ($apiWrapper.style.display === 'block') {
            $apiWrapper.style.display = 'none';
        } else {
            $apiWrapper.style.display = 'block';
            window.scrollTo(0, $apiWrapper.getBoundingClientRect().top);
        }
    });

    var paragraphs = document.querySelectorAll('.content p');
    var numParagraphs = paragraphs.length;
    for (var i = 0; i < numParagraphs; i++) {

        var clipboard = new ClipboardJS(paragraphs[i], {
            text: function (trigger) {
                return trigger.innerHTML;
            }
        });

        clipboard.on('success', function (e) {
            e.trigger.classList.remove('success');
            e.trigger.classList.add('success');
            setTimeout(function () {
                e.trigger.classList.remove('success');
            }, 2000);
        });

        paragraphs[i].addEventListener('click', function (e) {
            if (document.body.createTextRange) {
                const range = document.body.createTextRange();
                range.moveToElementText(e.currentTarget);
                range.select();
                document.execCommand('copy');
            } else if (window.getSelection) {
                const selection = window.getSelection();
                const range = document.createRange();
                range.selectNodeContents(e.currentTarget);
                selection.removeAllRanges();
                selection.addRange(range);
                document.execCommand('copy');
            } else {
                console.warn('Could not select text in node: Unsupported browser.');
            }
        });
    }
</script>

</body>
</html>
