"""Apply small fixes to the generated page head."""

import re
from urllib.parse import urljoin


ALTERNATE_LINK = re.compile(
    r'<link rel="alternate" href="([^"]+)" hreflang="([^"]+)">'
)
GLIGHTBOX_SCRIPT = re.compile(
    r'<script src="([^"]*assets/javascripts/glightbox\.min\.js)"></script>'
)
GLIGHTBOX_STYLESHEET = re.compile(
    r'<link href="([^"]*assets/stylesheets/glightbox\.min\.css)" rel="stylesheet">'
)
GLIGHTBOX_IMAGE_LINK = "element.setAttribute('href', img.src);"


def on_post_page(output, page, config, **kwargs):
    """Use absolute alternate URLs and keep the lightbox off the critical path."""
    base = page.canonical_url or config["site_url"]
    english_href = None

    def make_absolute(match):
        nonlocal english_href
        href = urljoin(base, match.group(1))
        lang = match.group(2)
        if lang == "en":
            english_href = href
        return f'<link rel="alternate" href="{href}" hreflang="{lang}">'

    output = ALTERNATE_LINK.sub(make_absolute, output)
    if english_href and 'hreflang="x-default"' not in output:
        english_tag = (
            f'<link rel="alternate" href="{english_href}" hreflang="en">'
        )
        output = output.replace(
            english_tag,
            english_tag
            + f'\n          <link rel="alternate" href="{english_href}" hreflang="x-default">',
            1,
        )

    script = GLIGHTBOX_SCRIPT.search(output)
    if script:
        output = GLIGHTBOX_SCRIPT.sub("", output, count=1)
        output = output.replace(
            '<script id="init-glightbox">',
            f'<script src="{script.group(1)}"></script>'
            '<script id="init-glightbox">',
            1,
        )
    output = output.replace(
        GLIGHTBOX_IMAGE_LINK,
        "if (!element.getAttribute('href')) { "
        "element.setAttribute('href', img.src); "
        "}",
    )
    return GLIGHTBOX_STYLESHEET.sub(
        r'<link href="\1" rel="stylesheet" media="print" '
        r"""onload="this.media='all'"><noscript><link href="\1" """
        r'rel="stylesheet"></noscript>',
        output,
    )
