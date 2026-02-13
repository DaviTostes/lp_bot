package utils

import "strings"

func TrimBullshit(output []byte) []byte {
	str := string(output)
	str = strings.TrimSpace(str)
	str = strings.TrimPrefix(str, "```js")
	str = strings.TrimPrefix(str, "```javascript")
	str = strings.TrimPrefix(str, "```tsx")
	str = strings.TrimPrefix(str, "```typescript")
	str = strings.TrimSuffix(str, "```")
	str = strings.TrimSpace(str)
	return []byte(str)
}
