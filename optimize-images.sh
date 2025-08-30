#!/bin/bash

# Image Optimization Script for Portfolio Website
# This script compresses images while preserving quality and resolution

set -e

echo "🚀 Starting image optimization..."

# Create backup directory
BACKUP_DIR="static/images-backup"
if [ ! -d "$BACKUP_DIR" ]; then
    echo "📦 Creating backup of original images..."
    cp -r static/images "$BACKUP_DIR"
fi

# Get initial size
INITIAL_SIZE=$(du -sh static/images | cut -f1)
echo "📊 Initial images size: $INITIAL_SIZE"

# Optimize PNG files using pngquant and ImageMagick
echo "🔧 Optimizing PNG files..."
find static/images -name "*.png" -type f | while read -r file; do
    echo "  Processing: $file"
    
    # First, use pngquant for color optimization (lossy but high quality)
    # --quality 80-95: High quality range
    # --skip-if-larger: Don't replace if result is larger
    if pngquant --quality 80-95 --skip-if-larger --strip --force --ext .png "$file" 2>/dev/null; then
        echo "    ✅ Compressed with pngquant"
    else
        echo "    ⚠️  pngquant skipped (would increase size)"
    fi
    
    # Then use ImageMagick for lossless optimization
    # -strip: Remove metadata
    # -interlace Plane: Progressive loading
    # -quality 95: High quality
    convert "$file" -strip -interlace Plane -quality 95 "$file.tmp"
    
    # Only replace if the new file is smaller
    if [ -f "$file.tmp" ] && [ $(stat -f%z "$file.tmp") -lt $(stat -f%z "$file") ]; then
        mv "$file.tmp" "$file"
        echo "    ✅ Further optimized with ImageMagick"
    else
        rm -f "$file.tmp"
        echo "    ⚠️  ImageMagick optimization skipped"
    fi
done

# Optimize JPEG files using jpegoptim
echo "🔧 Optimizing JPEG files..."
find static/images -name "*.jpg" -o -name "*.jpeg" -type f | while read -r file; do
    echo "  Processing: $file"
    
    # jpegoptim parameters:
    # --max=85: Maximum quality 85% (good balance of size/quality)
    # --strip-all: Remove all metadata
    # --preserve: Don't replace if result would be larger
    if jpegoptim --max=85 --strip-all --preserve "$file"; then
        echo "    ✅ Optimized JPEG"
    else
        echo "    ⚠️  JPEG optimization failed"
    fi
done

# Convert large PNG files to WebP for modern browsers (keeping originals)
echo "🔧 Creating WebP versions for large images..."
find static/images -name "*.png" -size +500k -type f | while read -r file; do
    webp_file="${file%.*}.webp"
    echo "  Creating WebP: $webp_file"
    
    # Create WebP version with high quality
    if convert "$file" -quality 85 "$webp_file"; then
        echo "    ✅ Created WebP version"
    else
        echo "    ⚠️  WebP conversion failed"
    fi
done

# Optimize the largest background image specifically
DESKTOP_BG="static/images/landing-page/desktop-bg.png"
if [ -f "$DESKTOP_BG" ]; then
    echo "🔧 Special optimization for desktop background..."
    
    # Create a high-quality JPEG version for faster loading
    DESKTOP_BG_JPG="${DESKTOP_BG%.*}.jpg"
    convert "$DESKTOP_BG" -quality 90 -strip "$DESKTOP_BG_JPG"
    echo "    ✅ Created JPEG version of desktop background"
    
    # Also create WebP
    DESKTOP_BG_WEBP="${DESKTOP_BG%.*}.webp"
    convert "$DESKTOP_BG" -quality 85 "$DESKTOP_BG_WEBP"
    echo "    ✅ Created WebP version of desktop background"
fi

# Final size calculation
FINAL_SIZE=$(du -sh static/images | cut -f1)
echo ""
echo "🎉 Optimization complete!"
echo "📊 Initial size: $INITIAL_SIZE"
echo "📊 Final size: $FINAL_SIZE"
echo "💾 Original images backed up to: $BACKUP_DIR"
echo ""
echo "🔍 Detailed size breakdown:"
du -sh static/images/* | sort -hr

echo ""
echo "💡 Tips for further optimization:"
echo "   - Consider using WebP images in your HTML templates"
echo "   - The desktop background has JPEG and WebP alternatives"
echo "   - All images have been optimized while preserving quality"
