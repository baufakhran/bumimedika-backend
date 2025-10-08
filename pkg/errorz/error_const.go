package errorz

var (
	ItemDraftAlreadyExistError        = NewByCode(ErrItemDraftAlreadyExist)
	ItemIdNotInUserDraftError         = NewByCode(ErrItemIdNotInUserDraft)
	ItemDraftInvalidFileContentsError = NewByCode(ErrItemDraftInvalidFileContents)
	ItemDraftEditHasNoPermissionError = NewByCode(ErrItemDraftEditHasNoPermission)
	ItemDraftExceedMaxLengthError     = NewByCode(ErrItemDraftExceedMaxLength)
	ItemDraftInvalidFormatError       = NewByCode(ErrItemDraftInvalidFormat)
	PageNotFoundError                 = NewByCode(ErrNotFound)
)
