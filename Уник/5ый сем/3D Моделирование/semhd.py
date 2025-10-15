import matplotlib.pyplot as plt
# номер 1
points = [(50, 50), (200, 80), (180, 200), (100, 220), (60, 150)] 
def fill_polygon(points):

    min_y = min(y for _, y in points)
    max_y = max(y for _, y in points)

    edges = []
    n = len(points)
    for i in range(n):
        p1 = points[i]
        p2 = points[(i + 1) % n]
        if p1[1] != p2[1]:  
            edges.append((p1, p2))

    filled_pixels = []

    for y in range(min_y, max_y + 1):
        intersections = []
        for p1, p2 in edges:
            if (p1[1] <= y < p2[1]) or (p2[1] <= y < p1[1]):
                x = p1[0] + (y - p1[1]) * (p2[0] - p1[0]) / (p2[1] - p1[1])
                intersections.append(x)
        intersections.sort()
        for i in range(0, len(intersections), 2):
            if i+1 < len(intersections):
                x_start = int(intersections[i])
                x_end = int(intersections[i+1])
                for x in range(x_start, x_end + 1):
                    filled_pixels.append((x, y))

    return filled_pixels

pixels = fill_polygon(points)

plt.figure(figsize=(6, 6))
plt.plot(*zip(*(points + [points[0]])), color='black')  
plt.scatter(*zip(*pixels), color='red', s=1)            
plt.gca().set_aspect('equal')
plt.show()

# номер 2
def get_edges(points):
    edges = []
    n = len(points)
    for i in range(n):
        p1 = points[i]
        p2 = points[(i + 1) % n]
        if p1[1] != p2[1]:
            edges.append((p1, p2))
    return edges

points = [(50, 50), (200, 80), (180, 200), (100, 220), (60, 150)]
non_horizontal_edges = get_edges(points)
print(non_horizontal_edges)


print("------------------------------\n")
# номер 3
def edge_horizontal_intersection(edge, y_line):
    (x1, y1), (x2, y2) = edge
    if y1 == y2:
        return None  
    if (y_line < min(y1, y2)) or (y_line > max(y1, y2)):
        return None 
    t = (y_line - y1) / (y2 - y1)
    x0 = t * (x2 - x1) + x1
    y0 = y_line
    return (x0, y0)

def get_intersections(points, y_line):
    edges = []
    n = len(points)
    for i in range(n):
        p1 = points[i]
        p2 = points[(i + 1) % n]
        if p1[1] != p2[1]:
            edges.append((p1, p2))
    intersections = []
    for edge in edges:
        pt = edge_horizontal_intersection(edge, y_line)
        if pt is not None:
            intersections.append(pt)
    return intersections

points = [(50, 50), (200, 80), (180, 200), (100, 220), (60, 150)]
y_line = 100
intersections = get_intersections(points, y_line)
print(f"Точки пересечения с y={y_line}: {intersections}")