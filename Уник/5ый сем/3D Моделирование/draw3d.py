from PIL import Image, ImageDraw
import sys
import math




def MdotM(M1, M2):
    M = []
    for i in range(4):
        M.append([0]*4)

    for i in range(4):
        for j in range(4):
            for k in range(4):
                M[i][j] += M1[i][k]*M2[k][j]
    return M


def MdotV(M, V):
    R = [0]*4
    for i in range(4):
        for j in range(4):
            R[i] += M[i][j]*V[j]
    return R


def parse_file(filename):
    data = {}
    data['vertices'] = []
    data['faces'] = []
    with open(filename) as f:
        for line in f:
            if line[0] == 'v':
                x, y, z = line[2:].split()
                data['vertices'].append((float(x), float(y), float(z)))
            elif line[0] == 'f' and line[1] == ' ':
                vec = line[2:].split()
                data['faces'].append(vec)

    return data

def draw_figure(data):
    img = Image.new('RGB', (200, 200))
    imdraw = ImageDraw.Draw(img)

    for v in data['vertices']:
        img.putpixel((int(v[0]), int(v[1])), (255, 0, 0))

    for f in data['faces']:
        size = len(f)
        for i in range(size):
            v1 = data['vertices'][int(f[i])-1]
            v2 = data['vertices'][int(f[(i+1)%size])-1]
            imdraw.line([(v1[0], v1[1]), (v2[0], v2[1])], width=1)

    return img


if __name__ == "__main__":
    filename = "pyramid.obj"
    if len(sys.argv)>1:
        filename = sys.argv[1]
    else:
        print("Usage: python3 main.py <objfile>")

    data = parse_file(filename)



    dx = 100
    dy = 100
    M = [[1, 0,  0, dx], 
         [0, 1, 0, dy],
         [0, 0, 1, 0],
         [0, 0, 0, 1]]

    s = 10
    S = [[s, 0, 0, 0],
         [0, s, 0, 0],
         [0, 0, s, 0],
         [0, 0, 0, 1]]

    V = [1, 2, 3, 1]

    Mi = MdotM(M, S)

    data1 = data.copy()
    for i in range(len(data['vertices'])):
        v = data['vertices'][i]
        extv = list(v)+[1]
        v = MdotV(Mi, extv)
        data1['vertices'][i] = v

    img = draw_figure(data1)

    img.save('figure.png')
