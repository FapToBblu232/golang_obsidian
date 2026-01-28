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


def isIn(face, points, x, y):
    for i in range(3):
        p0 = points[face[i]]
        p1 = points[face[(i+1) % 3]]

        p0o = (x - p0[0], y - p0[1])
        n = (p0[1] - p1[1], p1[0] - p0[0])

        if (p0o[0]*n[0] + p0o[1]*n[1]) > 0:
            return False
    return True


def plane_z(p1, p2, p3, x, y):
    x1,y1,z1 = p1
    x2,y2,z2 = p2
    x3,y3,z3 = p3

    A = (y2-y1)*(z3-z1) - (z2-z1)*(y3-y1)
    B = (z2-z1)*(x3-x1) - (x2-x1)*(z3-z1)
    C = (x2-x1)*(y3-y1) - (y2-y1)*(x3-x1)
    D = -A*x1 - B*y1 - C*z1

    if C == 0:
        return None

    return (-A*x - B*y - D) / C

vertices = [
    (0, 0, 0),
    (45, 0, 0),
    (45, 45, 0),
    (0, 45, 0),
    (0, 0, 90)
]

faces = [
    [0, 1, 2, 3],
    [1, 4, 0],
    [0, 3, 4]
]

colors = [
    (130, 130, 130),
    (200, 50, 50),
    (50, 200, 50)
]

S  = scale(2, 2, 2)
Ry = rotateY(math.radians(135))
Rx = rotateX(math.radians(70))
T  = translate(200, 200, 0)

M = MdotM(T, MdotM(Rx, MdotM(Ry, S)))

points = [transform(v, M) for v in vertices]

img = Image.new("RGB", (WIDTH, HEIGHT), (0,0,0))
zbuffer = [[float("inf")] * WIDTH for _ in range(HEIGHT)]

for fi, face in enumerate(faces):
    col = colors[fi]

    for t in range(1, len(face)-1):
        tri = [face[0], face[t], face[t+1]]

        xs = [points[i][0] for i in tri]
        ys = [points[i][1] for i in tri]

        minx = int(max(min(xs), 0))
        maxx = int(min(max(xs), WIDTH-1))
        miny = int(max(min(ys), 0))
        maxy = int(min(max(ys), HEIGHT-1))

        for y in range(miny, maxy+1):
            for x in range(minx, maxx+1):
                if not isIn(tri, points, x, y):
                    continue

                z = plane_z(points[tri[0]], points[tri[1]], points[tri[2]], x, y)
                if z is None:
                    continue

                if z < zbuffer[y][x]:
                    zbuffer[y][x] = z
                    img.putpixel((x, y), col)

img.show()
#img.save("figure_isIn.png")
