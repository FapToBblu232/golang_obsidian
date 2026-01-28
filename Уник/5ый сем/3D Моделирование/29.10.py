from math import sqrt
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


def podeli(p0, p1, t):
    return (p0[0] * (1 - t) + p1[0] * t,
            p0[1] * (1 - t) + p1[1] * t)

def distance_point_to_line_via_angle(px, py, x0, y0, x1, y1):
    ux = x1 - x0
    uy = y1 - y0
    vx = px - x0
    vy = py - y0

    u_len = sqrt(ux*ux + uy*uy)
    v_len = sqrt(vx*vx + vy*vy)
    if u_len == 0 or v_len == 0:
        return v_len

    dot = ux*vx + uy*vy
    cos_theta = dot / (u_len * v_len)
    if cos_theta > 1: cos_theta = 1
    if cos_theta < -1: cos_theta = -1

    sin_theta = sqrt(1 - cos_theta*cos_theta)
    return v_len * sin_theta

def bezier_recursive(img, p0, p1, p2, p3, d=2.0, color=(255,0,0), t=0.5):
    
    d1 = distance_point_to_line_via_angle(p1[0], p1[1], p0[0], p0[1], p3[0], p3[1])
    d2 = distance_point_to_line_via_angle(p2[0], p2[1], p0[0], p0[1], p3[0], p3[1])

    if d1 < d and d2 < d:
        draw_line(img, int(p0[0]), int(p0[1]), int(p3[0]), int(p3[1]), color)
        return
	
    p01 = podeli(p0, p1, t)
    p11 = podeli(p1, p2, t)
    p21 = podeli(p2, p3, t)
    p02 = podeli(p01, p11, t)
    p12 = podeli(p11, p21, t)
    p03 = podeli(p02, p12, t)

    bezier_recursive(img, p0, p01, p02, p03, d, color, t)
    bezier_recursive(img, p03, p12, p21, p3, d, color, t)



img = Image.new('RGB', (2000, 2000))
for i in range(10):
    for j in range(10):
        points = [(70 + i * 200, 10 + j * 200), (50 + i * 200, 100 + j * 200), (100 + i * 200, 150 + j * 200), (150 + i * 200, 100 + j * 200)]
        draw_line(img, points[0][0], points[0][1], points[1][0], points[1][1], (0, 255, 0))
        draw_line(img, points[1][0], points[1][1], points[2][0], points[2][1], (0, 255, 0))
        draw_line(img, points[2][0], points[2][1], points[3][0], points[3][1], (0, 255, 0))

        bezier_recursive(img, points[0], points[1], points[2], points[3], d=1.0, color=(255, 0, 0))
# bezier_recursive(img, points[0], points[1], points[2], points[3], d=1.0, color=(0, 255, 0), t=0.9);
img.show()