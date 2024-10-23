package application

type FileReader interface {
	FileReaderController() string

	FileWriterController() bool
}
