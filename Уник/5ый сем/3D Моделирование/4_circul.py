from PIL import Image

def draw_me(img, x0, y0, r, color):
    d = 2 - 2 * r
    x, y = 0, r
    while y >= 0:
        img.putpixel((x0 + x, y0 + y), color)
        img.putpixel((x0 + x, y0 - y), color)
        img.putpixel((x0 - x, y0 + y), color)
        img.putpixel((x0 - x, y0 - y), color)
        if (d < 0):
            error = 2 * d + 2 * y - 1
            if error <= 0:
                x += 1
                d = d + 2 * x + 1
                continue
            else:
                x += 1
                y -= 1
                d = d + 2 * x - 2 * y + 2
                continue 
        elif (d > 0):
            error = 2 * d - 2 * x - 1
            if error <= 0:
                x += 1
                y -= 1
                d = d + 2 * x - 2 * y + 2
                continue 
            else:
                y -= 1
                d = d - 2 * y + 1
                continue
        else:
            x += 1
            y -= 1
            d = d + 2 * x - 2 * y + 2



img = Image.new('RGB', (200, 200))
draw_me(img, 100, 100, 50, (255, 0, 0))
img.show()