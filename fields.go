package alpaca

// Address control field
func (a *Alpaca) Address(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// CKEditor control field
func (a *Alpaca) CKEditor(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Color control field
func (a *Alpaca) Color(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// ColorPicker control field
func (a *Alpaca) ColorPicker(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Country control field
func (a *Alpaca) Country(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Currency control field
func (a *Alpaca) Currency(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Date control field
func (a *Alpaca) Date(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// DateTime control field
func (a *Alpaca) DateTime(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Editor control field
func (a *Alpaca) Editor(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Email control field
func (a *Alpaca) Email(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Grid control field
func (a *Alpaca) Grid(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Image control field
func (a *Alpaca) Image(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Integer control field
func (a *Alpaca) Integer(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// IPv4 control field
func (a *Alpaca) IPv4(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// JSON control field
func (a *Alpaca) JSON(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Lowercase control field
func (a *Alpaca) Lowercase(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Map control field
func (a *Alpaca) Map(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// OptionTree control field
func (a *Alpaca) OptionTree(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Password control field
func (a *Alpaca) Password(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// PersonalName control field
func (a *Alpaca) PersonalName(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Phone control field
func (a *Alpaca) Phone(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// PickAColor control field
func (a *Alpaca) PickAColor(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Search control field
func (a *Alpaca) Search(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// State control field
func (a *Alpaca) State(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Summernote control field
func (a *Alpaca) Summernote(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Table control field
func (a *Alpaca) Table(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// TableRow control field
func (a *Alpaca) TableRow(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Tag control field
func (a *Alpaca) Tag(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Time control field
func (a *Alpaca) Time(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// TinyMCE control field
func (a *Alpaca) TinyMCE(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Token control field
func (a *Alpaca) Token(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Upload control field
func (a *Alpaca) Upload(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Uppercase control field
func (a *Alpaca) Uppercase(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// URL control field
func (a *Alpaca) URL(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Zipcode control field
func (a *Alpaca) Zipcode(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Any control field
func (a *Alpaca) Any(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Array container field
func (a *Alpaca) Array(f *Field) {
	f.IsContainerField = true

	maxItems := 1
	if f.Schema.Exists("maxItems") {
		maxItems = f.Schema.S("maxItems").Data().(int)
	}

	children, _ := f.Schema.S("items").ChildrenMap()
	childIndex := 0
	for key := range children {
		if childIndex < maxItems {
			a.ResolvePropertySchemaOptions(key+"["+string(childIndex)+"]", f)
		}
	}
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// File control field
func (a *Alpaca) File(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Hidden control field
func (a *Alpaca) Hidden(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Number control field
func (a *Alpaca) Number(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Object container field
func (a *Alpaca) Object(f *Field) {
	f.IsContainerField = true
	properties, err := f.Schema.S("properties").ChildrenMap()
	if err == nil {
		for key := range properties {
			a.ResolvePropertySchemaOptions(key, f)
		}
	}
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Text control field
func (a *Alpaca) Text(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// TextArea control field
func (a *Alpaca) TextArea(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Information container field
func (a *Alpaca) Information(f *Field) {
	f.IsContainerField = true
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Camera container field
func (a *Alpaca) Camera(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Repeatable container field
func (a *Alpaca) Repeatable(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Signature container field
func (a *Alpaca) Signature(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Checkbox container field
func (a *Alpaca) Checkbox(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Chooser container field
func (a *Alpaca) Chooser(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Radio container field
func (a *Alpaca) Radio(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

// Select container field
func (a *Alpaca) Select(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}
