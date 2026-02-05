#!/usr/bin/env python3
"""Generate intro.gif showcasing various themes."""
import subprocess
import sys
from pathlib import Path
from ansi2html import Ansi2HTMLConverter
from playwright.sync_api import sync_playwright
from PIL import Image
import shutil

PROJECT_DIR = Path(__file__).parent.parent
FRAMES_DIR = PROJECT_DIR / "assets" / "frames"
OUTPUT_GIF = PROJECT_DIR / "assets" / "intro.gif"

# Selected themes for the intro GIF (mix of categories)
SHOWCASE_THEMES = [
    "classic_framed",  # Default
    "cyberpunk",       # Sci-Fi
    "synthwave",       # Retro-Future
    "eva",             # Anime - Classic
    "dragonball",      # Anime - Classic
    "totoro",          # Ghibli
    "spirited",        # Ghibli
    "mecha",           # Anime Aesthetic
    "mahou",           # Anime Aesthetic
    "shonen",          # Anime Aesthetic
    "chibi",           # Anime Aesthetic
    "samurai",         # Anime Aesthetic
    "idol",            # Anime Aesthetic
    "spyfamily",       # Modern Anime
    "jujutsu",         # Modern Anime
    "matrix",          # Sci-Fi
    "htop",            # System Monitor
    "pixel",           # Retro
]

html_template = '''<!DOCTYPE html>
<html>
<head>
<style>
@import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500&display=swap');
* {{ margin: 0; padding: 0; box-sizing: border-box; }}
body {{
    background: #0d1117;
    padding: 16px 20px;
    font-family: 'JetBrains Mono', 'Monaco', 'Menlo', monospace;
    font-size: 13px;
    line-height: 1.5;
    display: inline-block;
}}
pre {{ margin: 0; white-space: pre; font-family: inherit; }}
</style>
</head>
<body><pre>{content}</pre></body>
</html>
'''

# Clean and create frames directory
if FRAMES_DIR.exists():
    shutil.rmtree(FRAMES_DIR)
FRAMES_DIR.mkdir(parents=True, exist_ok=True)

conv = Ansi2HTMLConverter(inline=True, dark_bg=True)

print("Generating frames...")
with sync_playwright() as p:
    browser = p.chromium.launch()
    page = browser.new_page()

    for i, theme in enumerate(SHOWCASE_THEMES):
        print(f"  Frame {i+1}/{len(SHOWCASE_THEMES)}: {theme}")
        result = subprocess.run(
            ["./statusline", "--preview", theme],
            capture_output=True, text=True, cwd=PROJECT_DIR
        )
        html_content = conv.convert(result.stdout, full=False)
        page.set_content(html_template.format(content=html_content))
        page.wait_for_timeout(200)
        page.locator("body").screenshot(path=str(FRAMES_DIR / f"frame_{i:03d}.png"))

    browser.close()

print("\nCreating GIF...")
# Load all frames
frames = []
for i in range(len(SHOWCASE_THEMES)):
    frame_path = FRAMES_DIR / f"frame_{i:03d}.png"
    img = Image.open(frame_path)
    # Convert to palette mode for better GIF quality
    img = img.convert('P', palette=Image.ADAPTIVE, colors=256)
    frames.append(img)

# Save as animated GIF (1500ms = 1.5 seconds per frame)
frames[0].save(
    OUTPUT_GIF,
    save_all=True,
    append_images=frames[1:],
    duration=1500,  # milliseconds per frame
    loop=0  # infinite loop
)

# Cleanup frames
shutil.rmtree(FRAMES_DIR)

print(f"Done! Created {OUTPUT_GIF}")
