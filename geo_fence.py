#!/usr/bin/env python
# -*- coding: utf-8 -*-


def geo_fence(x, y, points):
    count = 0
    x1, y1 = points[0]
    x1_part = (y1 > y) or ((x1 - x > 0) and (y1 == y)) # x1在哪一部分中
    x2, y2 = '', ''  # points[1]
    points.append((x1, y1))
    for point in points[1:]:
        x2, y2 = point
        x2_part = (y2 > y) or ((x2 > x) and (y2 == y)) # x2在哪一部分中
        if x2_part == x1_part:
            x1, y1 = x2, y2
            continue
        mul = (x1 - x)*(y2 - y) - (x2 - x)*(y1 - y)
        if mul > 0:  # 叉积大于0 逆时针
            count += 1
        elif mul < 0:
            count -= 1
        x1, y1 = x2, y2
        x1_part = x2_part
    if count == 2 or count == -2:
        return True
    else:
        return False


if __name__ == '__main__':
    points = [[116.239117,40.006047],[116.250103,39.835425],[116.483563,39.856513],[116.494549,39.995527]]
    x = 116.402538
    y = 39.922896
    for i in xrange(10000):
        geo_fence(x + i * 0.01, y + i * 0.01, points)
