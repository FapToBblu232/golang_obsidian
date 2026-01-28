from PIL import Image
import math

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
		
def draw_rect(img, points, color):
    for i in range(len(points)):
        current_point = points[i]
        next_point = points[(i + 1) % len(points)]
        draw_line(img, 
                current_point[0], current_point[1], 
                next_point[0], next_point[1], 
                color)


def apply_matrix(mat, point):
    x, y = point
    px = int(mat[0][0]*x + mat[0][1]*y + mat[0][2]*1)
    py = int(mat[1][0]*x + mat[1][1]*y + mat[1][2]*1)
    return (px, py)

def translate_matrix(dx, dy):
    return [
        [1, 0, dx],
        [0, 1, dy],
        [0, 0, 1]
    ]

def scale_matrix(sx, sy):
    return [
        [sx, 0,  0],
        [0,  sy, 0],
        [0,  0,  1]
    ]

def rotate_matrix(angle_rad):
    c = math.cos(angle_rad)
    s = math.sin(angle_rad)
    return [
        [ c, -s, 0],
        [ s,  c, 0],
        [ 0,  0, 1]
    ]


def transform_points(points, mat):
    return [apply_matrix(mat, p) for p in points]

def find_center(points):
    return (points[0][0] + points[2][0]) / 2, (points[0][1] + points[2][1]) / 2

def rotate_centre(points, angle_rad):
    centre = find_center(points)
    transf1 = translate_matrix(centre[0], centre[1])
    transf2 = translate_matrix(centre[0] * (-1), centre[1] * (-1))
    rotate = rotate_matrix(angle_rad)
    p1 = transform_points(points, transf2)
    p2 = transform_points(p1, rotate)
    return transform_points(p2, transf1)


img = Image.new('RGB', (200, 200))
points = [(0, 0), (1, 0), (1, 1), (0, 1)]
tr = translate_matrix(10,20)
rt = rotate_matrix(math.pi/6)
sc = scale_matrix(10, 15)
points1 = transform_points(points, sc)
points2 = transform_points(points1, tr)
points3 = transform_points(points2, rt)

draw_rect(img, points3, (255,0,0))
draw_rect(img, points2, (255,255,0))
draw_rect(img, points1, (255,0,255))
points4 = rotate_centre(points2, math.pi/6)
draw_rect(img, points, (255,255, 255))
draw_rect(img, points4, (255,255, 255))

img.show()
