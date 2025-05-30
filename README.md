# golang-design-pattern
- 学习一下go的设计模式

## 设计模式列表

### 创建型模式 (Creational Patterns)
1. **Factory Pattern (工厂模式)** - `factory/`
   - 根据输入参数创建不同类型的对象
   
2. **Builder Pattern (建造者模式)** - `builder/`
   - 逐步构建复杂对象
   
3. **Singleton Pattern (单例模式)** - `singleton/`
   - 确保一个类只有一个实例

### 结构型模式 (Structural Patterns)
4. **Adapter Pattern (适配器模式)** - `adapter/`
   - 让不兼容的接口能够协同工作
   
5. **Decorator Pattern (装饰器模式)** - `decorator/`
   - 动态地给对象添加新功能
   
6. **Facade Pattern (外观模式)** - `facade/`
   - 为复杂子系统提供简单接口
   
7. **Proxy Pattern (代理模式)** - `proxy/`
   - 为其他对象提供代理以控制对它的访问

### 行为型模式 (Behavioral Patterns)
8. **Observer Pattern (观察者模式)** - `observer/`
   - 定义对象间的一对多依赖关系
   
9. **Strategy Pattern (策略模式)** - `strategy/`
   - 定义算法族，使它们可以互相替换
   
10. **Command Pattern (命令模式)** - `command/`
    - 将请求封装成对象，支持撤销操作
    
11. **Template Method Pattern (模板方法模式)** - `template/`
    - 定义算法骨架，子类实现具体步骤
    
12. **State Pattern (状态模式)** - `state/`
    - 允许对象在内部状态改变时改变它的行为
    
13. **Chain of Responsibility Pattern (责任链模式)** - `chain/`
    - 将请求沿着处理者链传递，直到有处理者处理它
    
14. **Iterator Pattern (迭代器模式)** - `iterator/`
    - 提供一种方法顺序访问聚合对象中的各个元素

## 运行测试

```bash
# 运行所有测试
go test ./...

# 运行特定模式的测试
go test ./factory
go test ./builder
go test ./singleton
go test ./adapter
go test ./decorator
go test ./facade
go test ./proxy
go test ./observer
go test ./strategy
go test ./command
go test ./template
go test ./state
go test ./chain
go test ./iterator
```

## 设计模式说明

每个设计模式都包含：
- 实现代码 (`*.go`)
- 单元测试 (`*_test.go`)
- 实际使用示例

这些模式展示了如何在 Go 语言中优雅地解决常见的设计问题。