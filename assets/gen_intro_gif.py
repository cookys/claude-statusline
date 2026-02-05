#!/usr/bin/env python3
"""Generate a polished intro.gif with smooth transitions."""
import subprocess
from pathlib import Path
from ansi2html import Ansi2HTMLConverter
from playwright.sync_api import sync_playwright
from PIL import Image
import shutil

PROJECT_DIR = Path(__file__).parent.parent
FRAMES_DIR = PROJECT_DIR / "assets" / "frames"
OUTPUT_GIF = PROJECT_DIR / "assets" / "intro.gif"

# Curated showcase themes - variety of styles
SHOWCASE_THEMES = [
    ("classic_framed", "Default Theme"),
    ("cyberpunk", "Cyberpunk"),
    ("eva", "Evangelion"),
    ("totoro", "Totoro"),
    ("mecha", "Mecha HUD"),
    ("mahou", "Magical Girl"),
    ("samurai", "Samurai"),
    ("chibi", "Chibi"),
    ("synthwave", "Synthwave"),
    ("jujutsu", "Jujutsu Kaisen"),
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
    min-width: 800px;
}}
pre {{ margin: 0; white-space: pre; font-family: inherit; }}
.theme-label {{
    color: #8b949e;
    margin-bottom: 8px;
    font-size: 11px;
}}
</style>
</head>
<body>
<div class="theme-label">{label}</div>
<pre>{content}</pre>
</body>
</html>
'''

TARGET_SIZE = (850, 450)  # Consistent size for all frames

def resize_to_target(img):
    """Resize image to target size with padding."""
    img = img.convert('RGBA')
    # Create background
    bg = Image.new('RGBA', TARGET_SIZE, (13, 17, 23, 255))  # GitHub dark bg
    # Paste image centered or at top-left
    bg.paste(img, (0, 0))
    return bg

def create_crossfade_frames(img1, img2, steps=8):
    """Create crossfade transition frames between two images."""
    frames = []
    for i in range(steps):
        alpha = i / (steps - 1)
        blended = Image.blend(img1, img2, alpha)
        frames.append(blended.convert('P', palette=Image.ADAPTIVE, colors=256))
    return frames

# Clean and create frames directory
if FRAMES_DIR.exists():
    shutil.rmtree(FRAMES_DIR)
FRAMES_DIR.mkdir(parents=True, exist_ok=True)

conv = Ansi2HTMLConverter(inline=True, dark_bg=True)

print("Generating theme frames...")
theme_images = []

with sync_playwright() as p:
    browser = p.chromium.launch()
    page = browser.new_page()
    page.set_viewport_size({"width": 900, "height": 600})

    for theme, label in SHOWCASE_THEMES:
        print(f"  Capturing: {theme}")
        result = subprocess.run(
            ["./statusline", "--preview", theme],
            capture_output=True, text=True, cwd=PROJECT_DIR
        )
        html_content = conv.convert(result.stdout, full=False)
        full_label = f"Theme: {theme}"
        page.set_content(html_template.format(content=html_content, label=full_label))
        page.wait_for_timeout(300)

        # Save frame
        frame_path = FRAMES_DIR / f"{theme}.png"
        page.locator("body").screenshot(path=str(frame_path))

        img = Image.open(frame_path)
        img = resize_to_target(img)  # Normalize size
        theme_images.append(img)

    browser.close()

print("\nCreating animated GIF with transitions...")

all_frames = []
durations = []

for i, img in enumerate(theme_images):
    # Add the main frame (hold for 2 seconds)
    frame = img.convert('P', palette=Image.ADAPTIVE, colors=256)
    all_frames.append(frame)
    durations.append(2000)  # 2 seconds

    # Add crossfade to next theme
    if i < len(theme_images) - 1:
        next_img = theme_images[i + 1]
        fade_frames = create_crossfade_frames(img, next_img, steps=6)
        for ff in fade_frames[1:-1]:  # Skip first and last (they're the main frames)
            all_frames.append(ff)
            durations.append(80)  # Fast transition frames

# Loop back to first
fade_frames = create_crossfade_frames(theme_images[-1], theme_images[0], steps=6)
for ff in fade_frames[1:-1]:
    all_frames.append(ff)
    durations.append(80)

# Save animated GIF
all_frames[0].save(
    OUTPUT_GIF,
    save_all=True,
    append_images=all_frames[1:],
    duration=durations,
    loop=0,
    optimize=True
)

# Cleanup
shutil.rmtree(FRAMES_DIR)

file_size = OUTPUT_GIF.stat().st_size / 1024
print(f"\nDone! Created {OUTPUT_GIF}")
print(f"  Frames: {len(all_frames)}")
print(f"  Size: {file_size:.1f} KB")
