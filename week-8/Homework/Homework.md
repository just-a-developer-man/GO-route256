# Домашняя работа №8

## Caching
Реализовать кеширование в сервисе через Redis:
- [ ] Поднять Redis в docker-compose.
- [ ] Решить что кешировать и как, исходя из специфики системы, свой выбор обосновать (написать 1-2 предложения в README.md);
- [ ] Реализовать данное кеширование. Не нужно кешировать все, здесь достаточно одного типа данных.

## Sharding
Реализовать шардирование кеша:
- [ ] В docker-compose поднять два (или более) контейнеров Redis на разных портах (один контейнер будет представлять собой один шард);
- [ ] Решить как шардировать кешируемые данные и свой выбор обосновать (написать 1-2 предложения в README.md);
- [ ] Реализовать данное шардирование.

## System Design
💎 Спроектировать систему (написать примерно 1-2 абзаца в README.md)

Взять часть любой существующей системы (например, комментарии на YouTube, лента новостей в Twitter и т.д. что вам нравится) и описать, как бы вы ее спроектировали и почему именно так.
- [ ] Ответить как минимум на следующие вопросы:
- Кто потребители системы?
- Какие свойства критически важны для этой системы, а какие не очень?
- В зависимости от установленных приоритетных свойств описать:
    - Какие будут сервисы, на чем они будет написаны, что будут делать и как будут масштабированы?
    - Какие будут системы хранения данных, что будут хранить и как будут масштабированы?
    - Какие еще компоненты системы следует реализовать и почему (LB, RL и т.д.)?
- [ ] Нарисовать схему или описать, как это все будет друг с другом взаимодействовать.

## Profiling
💎 Улучшить код с помощью профайлинга (подкрепить результатами профайлинга "до" и "после" в README.md)
- [ ] Разобраться с профайлингом;
- [ ] Изучить скорость работы любого своего сервиса с помощью профайлинга;
- [ ] Найти узкие места в коде;
- [ ] Ускорить сервис, улучшив алгоритм в проблемных местах (можно тривиально, главное - понять, как работает профайлинг).

> Можно начать с профайлинга, чтобы после кеширования и шардирования естественным образом было видно ускорение. (Спасибо вам же за идею ;)