from abc import ABC, abstractmethod
from typing import Sequence
import itertools
import math
import matplotlib.pyplot as plt

class BaseNeuron(ABC):
    def __init__(self, initial_weights: Sequence[float], expected_output: Sequence[int], learning_speed: float = 0.3):
        if len(initial_weights) != 5:
            raise ValueError("Некорректный размер массива весов")
        self.weights = initial_weights

        if not (0 < learning_speed <= 1):
            raise ValueError("Скорость обучения не лежит в (0, 1]")
        self.learning_speed = learning_speed

        if len(expected_output) != 16:
            raise ValueError("Ошибка длины целевого вектора")
        self.expected_output = expected_output

        self.input_packages = list(itertools.product([0, 1], repeat=4))

    @abstractmethod
    def train(self):
        pass

    @abstractmethod
    def activation_function(self, input_values: Sequence[int]) -> int:
        pass

    # 
    def get_output(self) -> Sequence[int]:
        return [self.activation_function(package) for package in self.input_packages]

    def compute_err(self, generated_output: Sequence[int]) -> int:
        return sum(i != j for i, j in zip(generated_output, self.expected_output))

    @staticmethod
    def visualize_progress(number_of_Eras: int, err_list: Sequence[int], title: str):
        plt.title(title)
        plt.xlabel("Эпоха")
        plt.ylabel("Суммарная ошибка E")
        plt.grid(True)
        plt.plot(range(number_of_Eras + 1), err_list)
        plt.show()

    @staticmethod
    def bin_to_dec(binary_seq: Sequence[int]) -> int:
        return int(''.join(map(str, binary_seq)), 2)


class ThresholdNeiron(BaseNeuron):
    def updated_weights(self, input_package, err):
        self.weights = [wi + self.learning_speed * err * (1 if i == 0 else input_package[i - 1])
                     for i, wi in enumerate(self.weights)]

    def train(self):
        Eras_limit = 100
        err_log = []

        for era in range(Eras_limit):
            results = self.get_output()
            err = self.compute_err(results)
            print(f"Эпоха {era}: Y = {results}, W = {[round(float(v), 4) for v in self.weights]}, E = {err}")
            err_log.append(err)

            if err == 0:
                self.visualize_progress(era, err_log, "Зависимость суммарной ошибки НС от числа эпох")
                break

            #при помощи каждого набора проводим коррекцию (шаг обучения - тут)
            for idx, package in enumerate(self.input_packages):
                err = self.expected_output[idx] - self.activation_function(package)
                self.updated_weights(package, err)

    def activation_function(self, x_vec) -> int:
        net_value = sum(wi * xi for wi, xi in zip(self.weights[1:], x_vec)) + self.weights[0]
        return 1 if net_value >= 0 else 0


class TanhNeuron(BaseNeuron):
    def adjust_weights(self, package, delta, deriv):
        self.weights = [wi + self.learning_speed * delta * deriv * (1 if j == 0 else package[j - 1])
                     for j, wi in enumerate(self.weights)]

    def train(self, dataset=None, verbose=True) -> bool:
        if dataset is None:
            dataset = self.input_packages

        Eras_limit = 100
        err_log = []
        for era in range(Eras_limit):
            results = self.get_output()
            err = self.compute_err(results)
            err_log.append(err)
            if verbose:
                print(f"Эпоха {era}: Y = {results}, W = {[round(float(v), 3) for v in self.weights]}, E = {err}")

            if err == 0:
                if verbose:
                    self.visualize_progress(era, err_log, "Зависимость суммарной ошибки НС от числа эпох")
                return True

            for package in dataset:
                output = self.activation_function(package)
                delta = self.expected_output[self.bin_to_dec(package)] - output
                deriv = self.tahn_deriv(package)
                self.adjust_weights(package, delta, deriv)

        return False

    def selective_train(self):
        for count in range(2, 16):
            for subset in itertools.combinations(self.input_packages, count):
                self.weights = [0] * 5
                if self.train(subset, False):
                    print(f"Набор из {count} векторов: {subset}")
                    self.weights = [0] * 5
                    self.train(subset, True)
                    return

    def activation_function(self, x) -> int:
        return 1 if self.tanh_function(x) >= 0.5 else 0
    
    def tanh_function(self, x) -> float:
        net_value = sum(wi * xi for wi, xi in zip(self.weights[1:], x)) + self.weights[0]
        return float(0.5 * (math.tanh(net_value) + 1))

    def tahn_deriv(self, x) -> float:
        net_value = sum(wi * xi for wi, xi in zip(self.weights[1:], x)) + self.weights[0]
        try:
            return float(0.5 * ((1/math.cosh(net_value))**2))
        except ZeroDivisionError:
            return 1

target_values = [1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1]
print("№1 - Обучение с пороговой функцией активации")
print(f"БФ = {target_values}")
first_task = ThresholdNeiron([0] * 5, target_values)
first_task.train()
print()
print("№2 - Обучение с tanh-функцией")
print(f"БФ = {target_values}")
second_task = TanhNeuron([0] * 5, target_values)
second_task.train()
print()
print("№3 - Обучение на ограниченном наборе данных, tanh-фукнция")
print(f"БФ = {target_values}")
third_task = TanhNeuron([0] * 5, target_values)
third_task.selective_train()
print()