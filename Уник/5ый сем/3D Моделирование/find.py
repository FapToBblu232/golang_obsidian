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

def find_t(xa, ya, xb, yb, xi, yi, xj , yj):
	aby = yb - ya
	abx = xb - xa
	nx = yi - yj
	ny = xj - xi
	if (aby * ny + abx * nx) == 0:
		return None
	return (ny * (yi - ya) + nx * (xi - xa)) / (aby * ny + abx * nx)

img = Image.new('RGB', (200, 200))
points = [(70, 10), (50, 100), (100, 150), (150, 100), (130, 10)]

xa = 10
ya = 50
xb = 190
yb = 50

aby = yb - ya
abx = xb - xa

for i in range(len(points)):
    current_point = points[i]
    next_point = points[(i + 1) % len(points)]
    color = (255, 0, 0)
    nx = current_point[1] - next_point[1]
    ny = next_point[0] - current_point[0]
	
    if (abx * nx + ny * aby) < 0:
        color = (0, 0, 255)
	
    t = find_t(xa, ya, xb, yb, current_point[0], current_point[1], next_point[0], next_point[1])
    
    xk = t * (xb - xa) + xa
    yk = t * (yb - ya) + ya
	
    draw_line(img,
		current_point[0], current_point[1],
		xk, yk, color)
    
    draw_line(img, 
              current_point[0], current_point[1], 
              next_point[0], next_point[1], 
              color)
	

draw_line(img, xa, ya, xb, yb, (0, 255, 0))




    
img.show()