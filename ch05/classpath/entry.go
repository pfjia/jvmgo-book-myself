package classpath

import (
	"os"
	"strings"
)

const pathListSepartor = string(os.PathListSeparator)

type Entry interface {
	/**
	 * 读取class文件
	 * @param className class 文件的相对路径，如：java/lang/Object.class
	 */
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSepartor) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") || strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
