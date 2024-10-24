import cv2
import numpy as np

img1 = cv2.imread('random-image.png')
img2 = cv2.imread('encrypted1.png')

saida = cv2.bitwise_xor(img1, img2, mask = None)
saida2 = cv2.resize(saida, (540, 540))
cv2.imwrite('./Flag.png', saida2)
print("imagem gerada com sucesso")
