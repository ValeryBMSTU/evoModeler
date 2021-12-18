
import numpy as np
import random
import matplotlib.pyplot as plt

START_FROM = [
    'random',
    'form_one',
    'uniform'
]

ANT_SYS_MODIF = [
    'cycle',
    'density',
    'quantity'
]


class AntAlgorithm:

    def __init__(self,
                 a=1,  # коэффициент запаха
                 b=1,  # коэффициент расстояния
                 rho=0.5,  # коэффициент высыхания
                 Q=1,  # количество выпускаемого феромона
                 e=0,  # количество элитных муравьев
                 local_refresh=False,  # локальное обновление феромона
                 where_to_start='random',  # место начала движения муравьев
                 as_modif='density',  # модификация обновления феромона
                 random_seed=-1  # фиксация генератора случайных чисел
                 ):
        self.a = a
        self.b = b
        self.rho = rho
        self.Q = Q
        self.e = int(e)
        self.local_refresh = local_refresh
        if where_to_start in START_FROM:
            self.where_to_start = where_to_start
        else:
            raise ValueError('wrong where to start')
        if as_modif in ANT_SYS_MODIF:
            self.as_modif = as_modif
        else:
            raise ValueError('wrong ant-system modification name')
        if random_seed > 0:
            np.random.seed(random_seed)

    # START_FROM = {'random': randomly,
            #   'form_one': from_one,
            #   'uniform': uniform}

    def fit(self,
            L,  # Матрица смежности графа
            AGES=-1,  # количество поколений
            ANTS=-1,  # количество муравьев в поколении
            ph=-1  # начальное значение ферромона

            ):
        self.L = L
        self.CITIES = len(L)
        if AGES > 0:
            self.AGES = int(AGES)
        else:
            self.AGES = self.CITIES * 50
        if ANTS > 0:
            self.ANTS = int(ANTS)
        else:
            self.ANTS = self.CITIES
        if ph >= 0:
            self.ph = ph
        else:
            self.ph = self.Q/self.CITIES

        # инициализация матрицы "краткости" дуг графа
        rev_L = 1/L
        # инициализация матрицы феромонов
        tao = np.ones((self.CITIES, self.CITIES)) * self.ph

        self.BEST_DIST = float("inf")                # лучшая длина маршрута
        self.BEST_ROUTE = None                       # лучший маршрут
        # матрица маршрутов муравьев в одном поколении (номера узлов графа)
        antROUTE = np.zeros((self.ANTS, self.CITIES))
        # вектор длины маршрута муравьев в одном поколении
        antDIST = np.zeros(self.ANTS)
        # вектор лучших длин маршрутов в каждом поколении
        self.antBEST_DIST = np.zeros(self.AGES)
        self.antAVERAGE_DIST = np.zeros(self.AGES)

        # основной цикл алгоритма
        # ---------- начало освновного цикла ----------
        for age in range(self.AGES):
            antROUTE.fill(0)
            antDIST.fill(0)

            # ---------- начало цикла обхода графа муравьями ----------
            for k in range(self.ANTS):

                if self.where_to_start == 'random':
                    # начальное расположение муравья в графе (случайное)
                    antROUTE[k, 0] = np.random.randint(
                        low=0, high=self.CITIES-1, size=1)
                elif self.where_to_start == 'uniform':
                    # начальное расположение муравья в графе (равномерное)
                    antROUTE[k, 0] = k % self.CITIES
                elif self.where_to_start == 'from_one':
                    # начальное расположение муравья в графе (все с одного)
                    antROUTE[k, 0] = 1
                else:
                    assert 'wrong where to start'

                # ---------- начало обхода графа k-ым муравьем ----------
                for s in range(1, self.CITIES):
                    # текущее положение муравья
                    from_city = int(antROUTE[k, s-1])
                    P = (tao[from_city] ** self.a) * \
                        (rev_L[from_city] ** self.b)
                    # вероятность посещения уже посещенных городов = 0
                    for i in range(s):
                        P[int(antROUTE[k, i])] = 0

                    # вероятность выбора направления, сумма всех P = 1
                    assert (np.sum(P) > 0), \
                        "Division by zero. P = %s,"\
                        " \n tao = %s \n rev_L = %s" % (
                        P, tao[from_city], rev_L[from_city])
                    P = P / np.sum(P)
                    # выбираем направление
                    isNotChosen = True
                    while isNotChosen:
                        rand = np.random.rand()
                        for p, to in zip(P, list(range(self.CITIES))):
                            if p >= rand:
                                # записываем город №s в вектор k-ого муравья
                                antROUTE[k, s] = to
                                isNotChosen = False
                                break
                    # локальное обновление феромона
                    if self.local_refresh:
                        for s in range(self.CITIES):
                            city_to = int(antROUTE[k, s])
                            city_from = int(antROUTE[k, s-1])
                            if self.as_modif == 'cycle':
                                tao[city_from, city_to] = \
                                    tao[city_from, city_to] + \
                                    (self.Q / antDIST[k])
                            elif self.as_modif == 'density':
                                tao[city_from, city_to] = tao[city_from,
                                                              city_to] + self.Q
                            elif self.as_modif == 'quantity':
                                tao[city_from, city_to] = \
                                    tao[city_from, city_to] + \
                                    (self.Q / L[city_from, city_to])
                            else:
                                assert 'wrong ant-system modification name'
                            tao[city_to, city_from] = tao[city_from, city_to]
                # ---------- конец цила обхода графа ----------

                # вычисляем длину маршрута k-ого муравья
                for i in range(self.CITIES):
                    city_from = int(antROUTE[k, i-1])
                    city_to = int(antROUTE[k, i])
                    antDIST[k] += L[city_from, city_to]

                # сравниваем длину маршрута с лучшим показателем
                if antDIST[k] < self.BEST_DIST:
                    self.BEST_DIST = antDIST[k]
                    self.BEST_ROUTE = antROUTE[k]
            # ---------- конец цикла обхода графа муравьями ----------

            # ---------- обновление феромонов----------
            # высыхание по всем маршрутам (дугам графа)
            tao *= (1-self.rho)

            # цикл обновления феромона
            for k in range(self.ANTS):
                for s in range(self.CITIES):
                    city_to = int(antROUTE[k, s])
                    city_from = int(antROUTE[k, s-1])
                    if self.as_modif == 'cycle':
                        tao[city_from, city_to] = \
                            tao[city_from, city_to] + \
                            (self.Q / antDIST[k])
                    elif self.as_modif == 'density':
                        tao[city_from, city_to] = \
                            tao[city_from, city_to] + self.Q
                    elif self.as_modif == 'quantity':
                        tao[city_from, city_to] = \
                            tao[city_from, city_to] + \
                            (self.Q / L[city_from, city_to])
                    else:
                        assert 'wrong ant-system modification name'
                    tao[city_to, city_from] = tao[city_from, city_to]

            # проход элитных е-муравьев по лучшему маршруту
            if self.e > 0:
                for s in range(self.CITIES):
                    city_to = int(self.BEST_ROUTE[s])
                    city_from = int(self.BEST_ROUTE[s-1])
                    if self.as_modif == 'cycle':
                        tao[city_from, city_to] = \
                            tao[city_from, city_to] + \
                            ((self.Q * self.e) / self.BEST_DIST)
                    elif self.as_modif == 'density':
                        tao[city_from, city_to] = \
                            tao[city_from, city_to] + self.Q * self.e
                    elif self.as_modif == 'quantity':
                        tao[city_from, city_to] = \
                            tao[city_from, city_to] + \
                            ((self.Q * self.e) / L[city_from, city_to])
                    else:
                        assert 'wrong ant-system modification name'
                    tao[city_to, city_from] = tao[city_from, city_to]

            # ---------- конец обновления феромона ----------

            # конец поколения муравьев

            # сбор информации для графиков
            self.antBEST_DIST[age] = self.BEST_DIST
            self.antAVERAGE_DIST[age] = np.average(antDIST)

        return tao

    def get_best_route(self):
        return self.BEST_ROUTE

    def get_best_dist(self):
        return self.BEST_DIST

    def get_best_dists(self):
        return [age for age in range(self.AGES)], self.antBEST_DIST

    def get_average_dists(self):
        return [age for age in range(self.AGES)], self.antAVERAGE_DIST


# =============================================================================
L = np.array([[0, 2, 30, 9, 1],
              [4, 0, 47, 7, 7],
              [31, 33, 0, 33, 36],
              [20, 13, 16, 0, 28],
              [9, 36, 22, 22, 0]])
# =============================================================================

# =============================================================================
# L = np.random.randint(1, 50, size=(20, 20))
# np.fill_diagonal(L, 0)
# print(L)
# =============================================================================

# =============================================================================
# L = np.array([[0, 48, 16, 11, 28, 42, 49, 1, 24, 40, 14, 43, 12, 32, 39, 6, 42, 11, 39, 9],
#               [29, 0, 42, 19, 29, 41, 34, 29, 24, 2, 4, 15, 10, 17, 20, 37, 21, 15, 1, 19],
#               [29, 40, 0, 23, 29, 6, 18, 37, 36, 27, 2, 40, 42, 46, 22, 48, 41, 44, 15, 48],
#               [8, 23, 36, 0, 45, 8, 42, 49, 8, 25, 48, 2, 35, 9, 32, 46, 26, 12, 27, 31],
#               [32, 5, 26, 9, 0, 13, 16, 11, 34, 2, 10, 6, 6, 39, 45, 4, 19, 19, 26, 41],
#               [43, 6, 12, 18, 37, 0, 35, 43, 10, 5, 31, 22, 45, 49, 13, 38, 49, 35, 11, 14],
#               [39, 13, 29, 10, 19, 5, 0, 38, 44, 48, 18, 48, 49, 34, 26, 45, 11, 31, 33, 12],
#               [30, 29, 48, 30, 1, 12, 2, 0, 19, 28, 13, 38, 13, 4, 25, 5, 49, 38, 22, 42],
#               [5, 48, 9, 44, 34, 13, 9, 9, 0, 46, 1, 36, 4, 19, 20, 26, 25, 38, 48, 47],
#               [49, 21, 8, 48, 9, 34, 2, 20, 9, 0, 33, 38, 12, 9, 7, 46, 34, 16, 41, 1],
#               [40, 2, 20, 14, 35, 8, 6, 20, 33, 34, 0, 23, 8, 15, 14, 31, 28, 7, 8, 46],
#               [2, 8, 48, 40, 31, 31, 36, 24, 11, 45, 10, 0, 49, 40, 2, 27, 49, 42, 2, 4],
#               [40, 40, 12, 1, 11, 6, 46, 36, 36, 45, 47, 23, 0, 49, 4, 31, 9, 7, 23, 11],
#               [35, 18, 18, 7, 16, 37, 46, 36, 34, 45, 4, 12, 6, 0, 15, 30, 39, 42, 5, 47],
#               [46, 46, 38, 43, 40, 49, 37, 31, 16, 18, 15, 9, 27, 49, 0, 13, 27, 22, 19, 27],
#               [13, 44, 18, 44, 40, 23, 25, 9, 23, 13, 32, 22, 7, 28, 28, 0, 21, 26, 42, 2],
#               [3, 10, 46, 6, 23, 35, 47, 6, 40, 7, 25, 37, 26, 6, 6, 15, 0, 34, 15, 28],
#               [25, 9, 28, 12, 31, 46, 47, 45, 13, 46, 35, 35, 36, 20, 28, 13, 37, 0, 16, 24],
#               [31, 29, 32, 6, 29, 2, 6, 23, 14, 45, 43, 33, 32, 7, 22, 31, 35, 28, 0, 48],
#               [19, 42, 46, 15, 19, 19, 14, 32, 32, 19, 35, 39, 44, 7, 48, 20, 14, 14, 28, 0]])
# =============================================================================

# объявление коэффициентов

CITIES = len(L[0])
AGES = 50 * CITIES
ANTS = 20

a = 0.7  # коэффициент запаха
b = 1.5  # коэффициент расстояния
rho = 0.45  # коэффициент высыхания
Q = 120  # количество выпускаемого феромона
e = 5  # количество элитных муравьев

ph = Q / (CITIES)  # начальное значение феромона

# инициализация матрицы "краткости" дуг графа
rev_L = np.zeros((CITIES, CITIES))
for i in range(CITIES):
    for j in range(CITIES):
        if i != j:
            rev_L[i, j] = 1 / L[i, j]

# инициализация матрицы феромонов
tao = np.ones((CITIES, CITIES)) * ph

BEST_DIST = float("inf")  # лучшая длина маршрута
BEST_ROUTE = None  # лучший маршрут
antROUTE = np.zeros((ANTS, CITIES))  # матрица маршрутов муравьев в одном поколении (номера узлов графа)
antDIST = np.zeros(ANTS)  # вектор длины маршрута муравьев в одном поколении
antBEST_DIST = np.zeros(AGES)  # вектор лучших длин маршрутов в каждом поколении
antAVERAGE_DIST = np.zeros(AGES)

# основной цикл алгоритма
# ---------- начало освновного цикла ----------
for age in range(AGES):
    antROUTE.fill(0)
    antDIST.fill(0)

    # ---------- начало цикла обхода графа муравьями ----------
    for k in range(ANTS):

        # =============================================================================
        #         # начальное расположение муравья в графе (случайное)
        #         antROUTE[k, 0] = random.randint(0, CITIES-1)
        # =============================================================================

        # начальное расположение муравья в графе (равномерное)
        antROUTE[k, 0] = k % CITIES

        # =============================================================================
        #         # начальное расположение муравья в графе (все с одного)
        #         antROUTE[k, 0] = 1
        # =============================================================================

        # ---------- начало обхода графа k-ым муравьем ----------
        for s in range(1, CITIES):
            from_city = int(antROUTE[k, s - 1])  # текущее положение муравья
            P = (tao[from_city] ** a) * (rev_L[from_city] ** b)
            # вероятность посещения уже посещенных городов = 0
            for i in range(s):
                P[int(antROUTE[k, i])] = 0

            # вероятность выбора направления, сумма всех P = 1
            assert (np.sum(P) > 0), "Division by zero. P = %s, \n tao = %s \n rev_L = %s" % (
            P, tao[from_city], rev_L[from_city])
            P = P / np.sum(P)
            # выбираем направление
            isNotChosen = True
            while isNotChosen:
                rand = random.random()
                for p, to in zip(P, list(range(CITIES))):
                    if p >= rand:
                        antROUTE[k, s] = to  # записываем город №s в вектор k-ого муравья
                        isNotChosen = False
                        break
        # =============================================================================
        #             # локальное обновление феромона
        #             for s in range(CITIES):
        #                 city_to = int(antROUTE[k, s])
        #                 city_from = int(antROUTE[k, s-1])
        # #               tao[city_from, city_to] = tao[city_from, city_to] + (Q / antDIST[k]) # ant-cycle AntSystem
        #                 tao[city_from, city_to] = tao[city_from, city_to] + t # Ant-density AS
        #                 tao[city_to, city_from] = tao[city_from, city_to]
        # =============================================================================
        # ---------- конец цила обхода графа ----------

        # вычисляем длину маршрута k-ого муравья
        for i in range(CITIES):
            city_from = int(antROUTE[k, i - 1])
            city_to = int(antROUTE[k, i])
            antDIST[k] += L[city_from, city_to]

        # сравниваем длину маршрута с лучшим показателем
        if antDIST[k] < BEST_DIST:
            BEST_DIST = antDIST[k]
            BEST_ROUTE = antROUTE[k]
    # ---------- конец цикла обхода графа муравьями ----------

    # ---------- обновление феромонов----------
    # высыхание по всем маршрутам (дугам графа)
    tao *= (1 - rho)

    # цикл обновления феромона
    for k in range(ANTS):
        for s in range(CITIES):
            city_to = int(antROUTE[k, s])
            city_from = int(antROUTE[k, s - 1])
            tao[city_from, city_to] = tao[city_from, city_to] + (Q / antDIST[k])  # ant-cycle AntSystem
            #            tao[city_from, city_to] = tao[city_from, city_to] + Q # Ant-density AS
            tao[city_to, city_from] = tao[city_from, city_to]

    # проход элитных е-муравьев по лучшему маршруту
    for s in range(CITIES):
        city_to = int(BEST_ROUTE[s])
        city_from = int(BEST_ROUTE[s - 1])
        tao[city_from, city_to] = tao[city_from, city_to] + (e * Q / BEST_DIST)  # ant-cycle AntSystem
        tao[city_to, city_from] = tao[city_from, city_to]

    # ---------- конец обновления феромона ----------

    # конец поколения муравьев

    # сбор информации для графиков
    antBEST_DIST[age] = BEST_DIST
    antAVERAGE_DIST[age] = np.average(antDIST)

# ---------- конец основного цикла ----------

# выдача веса лучшего маршрута на выход
print(int(BEST_DIST))

x = list(range(AGES))
y = antBEST_DIST

fig, ax = plt.subplots()
plt.plot(x, y)
# =============================================================================
# for k in range(ANTS):
#     plt.plot(x, antALL_DIST[k])
# =============================================================================
plt.plot(x, antAVERAGE_DIST)
plt.grid(True)
plt.show()