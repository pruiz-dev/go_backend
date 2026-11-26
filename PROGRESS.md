# PROGRESS: Go Backend (subtracker)

## Текущее состояние
- **Активный этап:** Этап 1 — Рождение `subtracker`
- **Активный модуль:** `01_cli_calculator`
- **Текущий тикет:** `ticket_02` (Периоды оплаты и проверка условий (`if`, `switch`))
- **Статус:** Готово

## Что уже работает в `subtracker/`
- Инициализирован Go-модуль `subtracker` ([go.mod](file:///home/dima/Antigravity_projects/go_backend/subtracker/go.mod)).
- Консольная программа в [main.go](file:///home/dima/Antigravity_projects/go_backend/subtracker/main.go) и [subscription.go](file:///home/dima/Antigravity_projects/go_backend/subtracker/subscription.go): модель `Subscription`, валидация входных данных (`validateSubscription`), фабрика `newSubscription`, пересчёт периодов `weekly`/`monthly`/`yearly` (`calcSubscriptionCost`), обнуление неактивных подписок и табличный вывод.

## Замечания с ревью / на доработке
- Нет.

## Слабые места (для повторения через `/quiz`)
- Пока нет.

## Журнал последних сессий
- Репозиторий пересобран под формат Project-First + Offer-First. Готовы к первому тикету.
- Завершён `ticket_01` (`01_cli_calculator`): проект инициализирован, табличный вывод и расчёт стоимости работают.
- Завершён `ticket_02` (`01_cli_calculator`): добавлены периоды оплаты, валидация и расчёт через `if`/`switch`.