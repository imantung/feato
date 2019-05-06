# TYPFEAT

Typical Feature Toggles Library for Go.

- Global Access Instance (Singleton)
- Dynamic Handler (Realtime update)
- Avoid condition (No `if` statement)

## Usages

### New Feature
- [ ] `Name` string field
- [ ] `DefaultState` boolean field
- [ ] `State` boolean field


### Register & Run
- [ ] `Run(feat Feature, func() error) (err error)`
- [ ] `RunFeature(name string, default bool, func() error)`

### Enabling/Disabling
- [ ] `Enable(name string) error`
- [ ] `Disable(name string) error`

### Misc
- [ ] `listedFeatures() []Feature`: List of enabled features

### Recommendation

- Puts enabling/disabling function in single part of code

## Supporting Feature Storage

- [ ] Environment Variable

### Advanced Feature Toggle (Plan - TBD)

- [ ] Supporting embedded database - [bbolt](https://github.com/etcd-io/bbolt) as feature storage 
- [ ] Supporting external service integration e.g [flagr](https://github.com/checkr/flagr) as feature storage
- [ ] Add `tags` or `namespace` to categorize features and enabling/disabling by it 
- [ ] Add `target` to the feature to enable/disabling by run function parameter 
- [ ] Supporting Named Toggle


## References

- <https://medium.com/@dehora/feature-flags-smaller-better-faster-software-development-f2eab58df0f9>
- <https://martinfowler.com/articles/feature-toggles.html>
