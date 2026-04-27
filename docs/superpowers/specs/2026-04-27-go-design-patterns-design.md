# Go Design Patterns Showcase — Design Spec

## Purpose

A self-contained Go project demonstrating all three GoF design pattern categories (Creational, Structural, Behavioral) through idiomatic Go implementations. Built for personal learning — each pattern is independently runnable and testable.

## Architecture

### Layout

```
claude-test/
  go.mod
  cmd/
    creational/
      factory/main.go
      abstract-factory/main.go
      builder/main.go
      singleton/main.go
      prototype/main.go
    structural/
      adapter/main.go
      bridge/main.go
      composite/main.go
      decorator/main.go
      facade/main.go
      flyweight/main.go
      proxy/main.go
    behavioral/
      strategy/main.go
      observer/main.go
      command/main.go
      chain/main.go
      state/main.go
      template/main.go
  internal/
    creational/
      factory/factory.go, factory_test.go
      abstract_factory/abstract_factory.go, abstract_factory_test.go
      builder/builder.go, builder_test.go
      singleton/singleton.go, singleton_test.go
      prototype/prototype.go, prototype_test.go
    structural/
      adapter/adapter.go, adapter_test.go
      bridge/bridge.go, bridge_test.go
      composite/composite.go, composite_test.go
      decorator/decorator.go, decorator_test.go
      facade/facade.go, facade_test.go
      flyweight/flyweight.go, flyweight_test.go
      proxy/proxy.go, proxy_test.go
    behavioral/
      strategy/strategy.go, strategy_test.go
      observer/observer.go, observer_test.go
      command/command.go, command_test.go
      chain/chain.go, chain_test.go
      state/state.go, state_test.go
      template/template.go, template_test.go
```

### Roles

- **`cmd/<category>/<pattern>/main.go`** — thin entry point that imports from `internal/`, creates instances, and prints a demo to stdout. Each prints a header line: `=== Pattern Name Pattern ===`.
- **`internal/<category>/<pattern>/`** — types, interfaces, and logic for the pattern.
- **`*_test.go`** — exercises the pattern. Table-driven tests where the pattern has multiple variants; single-case tests for patterns with one flow.

### Running

- Single pattern demo: `go run ./cmd/creational/factory`
- All tests: `go test ./internal/...`
- Each pattern is fully independent — no cross-pattern imports.

## Pattern Inventory

### Creational (5)

| Pattern | Domain Example | Key Go Concept |
|---------|---------------|----------------|
| Factory Method | `NotificationFactory` — Email, SMS, Push notifications | Interface return types |
| Abstract Factory | `UIFactory` — Light/Dark theme widget families | Nested interfaces |
| Builder | `QueryBuilder` — constructs SQL queries step by step | Method chaining, functional options |
| Singleton | `ConfigManager` — single app config instance | `sync.Once` |
| Prototype | `DocumentTemplate` — cloneable document templates | Deep copy via interface |

### Structural (7)

| Pattern | Domain Example | Key Go Concept |
|---------|---------------|----------------|
| Adapter | `XMLToJSONAdapter` — converts XML data source to JSON interface | Interface wrapping |
| Bridge | `Renderer` x `Shape` — decouple rendering from shapes | Composition of interfaces |
| Composite | `FileSystem` — files and directories as a tree | Recursive interface |
| Decorator | `DataSource` — add compression, encryption layers | Interface wrapping + delegation |
| Facade | `MediaConverter` — simplifies ffmpeg-like subsystems | Struct aggregating subsystems |
| Flyweight | `CharacterRenderer` — shared glyph state for text rendering | Pooling with `map` |
| Proxy | `CachedWeatherService` — caching proxy for an API | Same interface, controlled access |

### Behavioral (6)

| Pattern | Domain Example | Key Go Concept |
|---------|---------------|----------------|
| Strategy | `Sorter` — pluggable sort algorithms | Function types as strategies |
| Observer | `EventBus` — pub/sub for app events | Channels or callback slices |
| Command | `TextEditor` — undo/redo command history | Interface with Execute/Undo |
| Chain of Responsibility | `HTTPMiddleware` — auth, logging, rate-limit chain | Linked handlers |
| State | `VendingMachine` — state-driven behavior | Interface per state |
| Template Method | `DataExporter` — CSV/JSON/XML export with shared pipeline | Embedded struct + interface |

## Per-Pattern Structure

Each pattern follows a consistent internal structure:

1. Define interfaces first — the contracts the pattern establishes.
2. Concrete types implement those interfaces.
3. A constructor function (e.g., `NewFactory(...)`) serves as the entry point.
4. No global state (except Singleton, by definition).

## Conventions

- **No external dependencies** — stdlib only, `go.mod` stays clean.
- **Consistent naming** — package name matches the pattern in snake_case (e.g., `abstract_factory`); `cmd/` directories use hyphens for readability (e.g., `abstract-factory`).
- **Errors returned, not panicked** — patterns that can fail return `error`.
- **No comments explaining GoF theory** — code and test names make the pattern self-evident.

## Testing Strategy

- Table-driven tests where the pattern has multiple variants (Factory, Strategy).
- Single-case tests for patterns with one flow (Singleton, Facade).
- Assert on behavior, not internal state.
- Tests verify the structural contract: Decorator preserves the interface, Observer delivers to all subscribers, Command supports undo, etc.

## Out of Scope

- No external dependencies or frameworks.
- No cross-pattern composition or shared domain model.
- No GoF theory documentation — this is a code showcase, not a textbook.
- No CI/CD, Docker, or deployment configuration.
