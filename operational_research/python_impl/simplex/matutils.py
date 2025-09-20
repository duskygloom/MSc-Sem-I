import numpy as np


def concat_mat_horiz(a: np.ndarray, b: np.ndarray) -> np.ndarray:
    if a.ndim != 2 or b.ndim != 2:
        raise Exception("EXPECTED 2D ARRAYS")
    if a.shape[0] != b.shape[0]:  # unequal rows
        raise Exception("BOTH MATRICES SHOULD HAVE EQUAL ROWS")
    c = np.zeros((a.shape[0], a.shape[1] + b.shape[1]))
    for i in range(a.shape[0]):
        for j in range(a.shape[1]):
            c[i, j] = a[i, j]
        for j in range(b.shape[1]):
            c[i, a.shape[1] + j] = b[i, j]
    return c


def concat_mat_vert(a: np.ndarray, b: np.ndarray) -> np.ndarray:
    if a.ndim != 2 or b.ndim != 2:
        raise Exception("EXPECTED 2D ARRAY")
    if a.shape[1] != b.shape[1]:  # unequal rows
        raise Exception("BOTH MATRICES SHOULD HAVE EQUAL COLS")
    c = np.zeros((a.shape[0] + b.shape[0], a.shape[1]))
    for i in range(a.shape[0]):
        for j in range(a.shape[1]):
            c[i, j] = a[i, j]
    for i in range(b.shape[0]):
        for j in range(a.shape[1]):
            c[a.shape[0] + i, j] = b[i, j]
    return c


def left_pad_mat(m: np.ndarray, padding: int) -> np.ndarray:
    if m.ndim != 2:
        raise Exception("EXPECTED 2D ARRAY")
    pad_mat = np.zeros((m.shape[0], padding), dtype=float)
    return concat_mat_horiz(pad_mat, m)
