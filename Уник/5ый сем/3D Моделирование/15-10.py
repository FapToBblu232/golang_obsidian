from PIL import Image

def draw_rect(img, points, color):
    for i in range(len(points)):
        current_point = points[i]
        next_point = points[(i + 1) % len(points)]
        draw_line(img, 
                current_point[0], current_point[1], 
                next_point[0], next_point[1], 
                color)

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

def get_side_array(points):
	answ = []
	for i in range(len(points)):
		current = points[i]
		next = points[(i + 1) % len(points)]
		if current[1] != next[1]:
			answ.append((current, next))
	return tuple(answ)

def find_t(side, y):
	return (y - side[0][1]) / (side[1][1] - side[0][1])
	
def get_intersection(side, y):
	t = find_t(side, y)
	x_0 = int(t * side[1][0] + (1 - t) * side[0][0])
	return (x_0, y)

def get_intersection_array(sides, y):
	answ = []
	for i in range(len(sides)):
		if y < min(sides[i][0][1], sides[i][1][1]) or y > max(sides[i][0][1], sides[i][1][1]):
			continue
		temp = get_intersection(sides[i], y)
		if temp not in answ:
			answ.append(temp)
	answ.sort(key=lambda point: point[0])
	return answ

def fill(img, inters, color):
	for i in range(0, len(inters) - 1, 2):
		draw_line(img, inters[i][0], inters[i][1], inters[i + 1][0], inters[i + 1][1], color)

def fill_xor(img, y, color):
    xor = 0
    color = img.getpixel((0, y))
    for i in range(1, 200, 1):
        if color != img.getpixel((i, y)):
            xor = (xor + 1) % 2
        if xor == 1:
            img.putpixel((i, y), (255, 0, 0))
		
		

img = Image.new('RGB', (200, 200))

# points = [(70, 10), (50, 100), (100, 150), (150, 100), (130, 10)]
points = [(70, 10), (100, 10), (100, 150), (70, 150)]
maxi = 0
mini = 200
for i in range(len(points)):
	if points[i][1] > maxi:
		maxi = points[i][1]
	if points[i][1] < mini:
		mini = points[i][1]
draw_rect(img, points, (255, 0,0))
sides = get_side_array(points)
for i in range(mini, maxi):
	inters = get_intersection_array(sides, i)
	print(inters)
	fill_xor(img, i, (255, 0, 0))
	

img.show()