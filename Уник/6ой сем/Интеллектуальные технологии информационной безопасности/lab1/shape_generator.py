import random
import math
from PIL import Image, ImageDraw

# --------------------------
# Настройки
# --------------------------
WIDTH, HEIGHT = 900, 700
BACKGROUND = (255, 255, 255)

COLORS = {
    "yellow": (255, 204, 0),
    "blue": (0, 102, 204),
    "red": (220, 0, 0)
}

SHAPES = ["star", "rounded_square", "rounded_triangle", "rounded_pentagon"]

MIN_SIZE = 40
MAX_SIZE = 100
MAX_ATTEMPTS = 5000

# --------------------------
# Проверка пересечения
# --------------------------
def bounding_box_overlap(box1, box2):
    return not (
        box1[2] < box2[0] or
        box1[0] > box2[2] or
        box1[3] < box2[1] or
        box1[1] > box2[3]
    )

# --------------------------
# Генерация фигур
# --------------------------
def generate_regular_polygon(cx, cy, radius, sides, rotation=0):
    points = []
    for i in range(sides):
        angle = 2 * math.pi * i / sides + rotation
        x = cx + radius * math.cos(angle)
        y = cy + radius * math.sin(angle)
        points.append((x, y))
    return points

def generate_star(cx, cy, radius):
    points = []
    for i in range(10):
        angle = math.pi * i / 5
        r = radius if i % 2 == 0 else radius / 2
        x = cx + r * math.cos(angle)
        y = cy + r * math.sin(angle)
        points.append((x, y))
    return points

# --------------------------
# Рисование скруглённого многоугольника
# --------------------------
def draw_rounded_polygon(draw, points, radius, fill):
    """
    Рисует многоугольник с закруглёнными углами через Pillow.
    """
    n = len(points)
    # рисуем линии с дугами на вершинах
    for i in range(n):
        p0 = points[i - 1]
        p1 = points[i]
        p2 = points[(i + 1) % n]

        # направления сторон
        dx1, dy1 = p0[0] - p1[0], p0[1] - p1[1]
        len1 = math.hypot(dx1, dy1)
        ux1, uy1 = dx1 / len1, dy1 / len1

        dx2, dy2 = p2[0] - p1[0], p2[1] - p1[1]
        len2 = math.hypot(dx2, dy2)
        ux2, uy2 = dx2 / len2, dy2 / len2

        # точки начала и конца дуги
        p1a = (p1[0] + ux1 * radius, p1[1] + uy1 * radius)
        p1b = (p1[0] + ux2 * radius, p1[1] + uy2 * radius)

        # соединяем линии
        draw.line([p1a, p1b], fill=fill, width=radius*2)

    # заливаем центр
    draw.polygon(points, fill=fill)

# --------------------------
# Создаём холст
# --------------------------
image = Image.new("RGB", (WIDTH, HEIGHT), BACKGROUND)
draw = ImageDraw.Draw(image)

placed_boxes = []

# Перемножение фигур и цветов
all_combinations = [(shape, color) for shape in SHAPES for color in COLORS.values()]

for shape, color in all_combinations:

    for _ in range(MAX_ATTEMPTS):
        size = random.randint(MIN_SIZE, MAX_SIZE)
        cx = random.randint(size, WIDTH - size)
        cy = random.randint(size, HEIGHT - size)

        box = (cx - size, cy - size, cx + size, cy + size)

        # проверка пересечения
        if any(bounding_box_overlap(box, existing) for existing in placed_boxes):
            continue

        if shape == "rounded_square":
            draw.rounded_rectangle(box, radius=size//4, fill=color)

        elif shape == "rounded_triangle":
            points = generate_regular_polygon(cx, cy, size, 3, rotation=random.random())
            draw_rounded_polygon(draw, points, radius=size//4, fill=color)

        elif shape == "rounded_pentagon":
            points = generate_regular_polygon(cx, cy, size, 5, rotation=random.random())
            draw_rounded_polygon(draw, points, radius=size//5, fill=color)

        elif shape == "star":
            points = generate_star(cx, cy, size)
            draw.polygon(points, fill=color)

        placed_boxes.append(box)
        break

# --------------------------
# Сохраняем
# --------------------------
image.save("output.png")
print("Готово! Сохранено в output.png")