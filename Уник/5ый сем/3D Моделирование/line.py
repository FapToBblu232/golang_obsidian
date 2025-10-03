from PIL import Image

def draw_line(img, x0, y0, x1, y1, color):
	dx = x1 - x0
	dy = y1 - y0
	
	sign_x = 1 if dx>0 else -1 if dx<0 else 0
	sign_y = 1 if dy>0 else -1 if dy<0 else 0
	
	if dx < 0: dx = -dx
	if dy < 0: dy = -dy
	
	if dx > dy:
		pdx, pdy = sign_x, 0
		es, el = dy, dx
	else:
		pdx, pdy = 0, sign_y
		es, el = dx, dy
	
	x, y = x0, y0
	e = 0
	img.putpixel((x, y), color)
	
	for i in range(el):
		e += 2 * es
		if e > el:
			e -= 2 * el
			x += sign_x
			y += sign_y
		else:
			x += pdx
			y += pdy
		img.putpixel((x, y), color)

# def draw_me(img, x0, y0, x1, y1, color):
	
# 	e = 0
# 	y = y0
# 	for i in range(abs(x1 - x0)):
# 		x = x0 + k2 * i
# 		e += k1 * 2 * abs(y1 - y0)
# 		if e > (x1 - x0):
# 			y += k1 * 1
# 			e -= 2 * abs(x1 - x0)
# 		img.putpixel((x, y), color)



img = Image.new('RGB', (200, 200))
draw_line(img, 100, 100, 150, 10, (255, 0, 0))
draw_line(img, 100, 100, 150, 70, (255, 0, 0))
draw_line(img, 100, 100, 150, 130, (255, 0, 0))
draw_line(img, 100, 100, 150, 190, (255, 0, 0))
draw_line(img, 100, 100, 50, 10, (255, 0, 0))
draw_line(img, 100, 100, 50, 70, (255, 0, 0))
draw_line(img, 100, 100, 50, 130, (255, 0, 0))
draw_line(img, 100, 100, 50, 190, (255, 0, 0))
img.show()