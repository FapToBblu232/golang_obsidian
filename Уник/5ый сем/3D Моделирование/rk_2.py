from PIL import Image
import math

WIDTH, HEIGHT = 400, 400

def MdotM(A, B):
    R = [[0]*4 for _ in range(4)]
    for i in range(4):
        for j in range(4):
            for k in range(4):
                R[i][j] += A[i][k] * B[k][j]
    return R

def MdotV(M, v):
    r = [0]*4
    for i in range(4):
        for j in range(4):
            r[i] += M[i][j] * v[j]
    return r

def translate(dx, dy, dz):
    return [
        [1, 0, 0, dx],
        [0, 1, 0, dy],
        [0, 0, 1, dz],
        [0, 0, 0, 1]
    ]

def scale(sx, sy, sz):
    return [
        [sx, 0,  0,  0],
        [0,  sy, 0,  0],
        [0,  0,  sz, 0],
        [0,  0,  0,  1]
    ]

def rotateX(a):
    c, s = math.cos(a), math.sin(a)
    return [
        [1, 0, 0, 0],
        [0, c,-s, 0],
        [0, s, c, 0],
        [0, 0, 0, 1]
    ]

def rotateY(a):
    c, s = math.cos(a), math.sin(a)
    return [
        [ c, 0, s, 0],
        [ 0, 1, 0, 0],
        [-s, 0, c, 0],
        [ 0, 0, 0, 1]
    ]

def transform(p, M):
    x, y, z = p
    r = MdotV(M, [x, y, z, 1])
    return (r[0], r[1], r[2])

def project(p):
    return int(p[0]), int(p[1]), p[2]

def barycentric(px, py, a, b, c):
    det = (b[1]-c[1])*(a[0]-c[0]) + (c[0]-b[0])*(a[1]-c[1])
    if det == 0:
        return -1, -1, -1

    w1 = ((b[1]-c[1])*(px-c[0]) + (c[0]-b[0])*(py-c[1])) / det
    w2 = ((c[1]-a[1])*(px-c[0]) + (a[0]-c[0])*(py-c[1])) / det
    w3 = 1 - w1 - w2
    return w1, w2, w3

def draw_triangle(img, zbuf, v1, v2, v3, color):
    x1, y1, z1 = project(v1)
    x2, y2, z2 = project(v2)
    x3, y3, z3 = project(v3)

    minx = max(min(x1, x2, x3), 0)
    maxx = min(max(x1, x2, x3), WIDTH-1)
    miny = max(min(y1, y2, y3), 0)
    maxy = min(max(y1, y2, y3), HEIGHT-1)

    for y in range(miny, maxy+1):
        for x in range(minx, maxx+1):
            w1, w2, w3 = barycentric(x, y, (x1,y1), (x2,y2), (x3,y3))
            if w1 < 0 or w2 < 0 or w3 < 0:
                continue

            z = w1*z1 + w2*z2 + w3*z3
            if z < zbuf[y][x]:
                zbuf[y][x] = z
                img.putpixel((x, y), color)


vertices = [
    (0, 0, 0),
    (45, 0, 0),
    (45, 45, 0),
    (0, 45, 0),
    (0, 0, -90)
]

faces = [
    [0, 1, 2, 3],
    [1, 2, 4],
    [2, 3, 4]
]

S = scale(2, 2, 2)
Ry = rotateY(math.radians(135))
Rx = rotateX(math.radians(10))
T = translate(200, 200, 0)

M = MdotM(T, MdotM(Rx, MdotM(Ry, S)))

verts = [transform(v, M) for v in vertices]

# Изображение и Z-buffer
img = Image.new("RGB", (WIDTH, HEIGHT), (0,0,0))
zbuffer = [[float("inf")]*WIDTH for _ in range(HEIGHT)]

colors = [
    (120, 120, 120),   # основание
    (200, 50, 50),
    (50, 200, 50)
]

for i, face in enumerate(faces):
    vs = [verts[idx] for idx in face]
    for j in range(1, len(vs)-1):
        draw_triangle(img, zbuffer, vs[0], vs[j], vs[j+1], colors[i])

img.show()
#img.save("figure_3_faces.png")
