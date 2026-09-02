"""Put a copy of the sitemap inside every language folder.

The theme reads the sitemap that sits next to each language link so the
language button can open the same page in the other language. Only the site
root gets a sitemap, so the language folder answered with "not found".

The priority keeps this hook running after the translated pages are written.
"""

import shutil
from pathlib import Path

from mkdocs.plugins import event_priority

SITEMAP_FILES = ("sitemap.xml", "sitemap.xml.gz")


@event_priority(-150)
def on_post_build(config, **kwargs):
    plugin = config["plugins"].get("i18n")
    if plugin is None:
        return

    site_dir = Path(config["site_dir"])
    for language in plugin.config.languages:
        if language.default or not language.build:
            continue

        folder = site_dir / language.locale
        if not folder.is_dir():
            continue

        for name in SITEMAP_FILES:
            source = site_dir / name
            if source.is_file():
                shutil.copyfile(source, folder / name)
