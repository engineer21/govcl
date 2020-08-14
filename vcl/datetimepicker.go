
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    "time"
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TDateTimePicker struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewDateTimePicker(owner IComponent) *TDateTimePicker {
    d := new(TDateTimePicker)
    d.instance = DateTimePicker_Create(CheckPtr(owner))
    d.ptr = unsafe.Pointer(d.instance)
    return d
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsDateTimePicker(obj interface{}) *TDateTimePicker {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TDateTimePicker{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsDateTimePicker.
func DateTimePickerFromInst(inst uintptr) *TDateTimePicker {
    return AsDateTimePicker(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsDateTimePicker.
func DateTimePickerFromObj(obj IObject) *TDateTimePicker {
    return AsDateTimePicker(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsDateTimePicker.
func DateTimePickerFromUnsafePointer(ptr unsafe.Pointer) *TDateTimePicker {
    return AsDateTimePicker(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (d *TDateTimePicker) Free() {
    if d.instance != 0 {
        DateTimePicker_Free(d.instance)
        d.instance, d.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (d *TDateTimePicker) Instance() uintptr {
    return d.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (d *TDateTimePicker) UnsafeAddr() unsafe.Pointer {
    return d.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (d *TDateTimePicker) IsValid() bool {
    return d.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (d *TDateTimePicker) Is() TIs {
    return TIs(d.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (d *TDateTimePicker) As() TAs {
//    return TAs(d.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TDateTimePickerClass() TClass {
    return DateTimePicker_StaticClassType()
}

func (d *TDateTimePicker) DateIsNull() bool {
    return DateTimePicker_DateIsNull(d.instance)
}

func (d *TDateTimePicker) SelectDate() {
    DateTimePicker_SelectDate(d.instance)
}

func (d *TDateTimePicker) SelectTime() {
    DateTimePicker_SelectTime(d.instance)
}

// 是否可以获得焦点。
func (d *TDateTimePicker) CanFocus() bool {
    return DateTimePicker_CanFocus(d.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (d *TDateTimePicker) ContainsControl(Control IControl) bool {
    return DateTimePicker_ContainsControl(d.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (d *TDateTimePicker) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(DateTimePicker_ControlAtPos(d.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (d *TDateTimePicker) DisableAlign() {
    DateTimePicker_DisableAlign(d.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (d *TDateTimePicker) EnableAlign() {
    DateTimePicker_EnableAlign(d.instance)
}

// 查找子控件。
//
// Find sub controls.
func (d *TDateTimePicker) FindChildControl(ControlName string) *TControl {
    return AsControl(DateTimePicker_FindChildControl(d.instance, ControlName))
}

func (d *TDateTimePicker) FlipChildren(AllLevels bool) {
    DateTimePicker_FlipChildren(d.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (d *TDateTimePicker) Focused() bool {
    return DateTimePicker_Focused(d.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (d *TDateTimePicker) HandleAllocated() bool {
    return DateTimePicker_HandleAllocated(d.instance)
}

// 插入一个控件。
//
// Insert a control.
func (d *TDateTimePicker) InsertControl(AControl IControl) {
    DateTimePicker_InsertControl(d.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (d *TDateTimePicker) Invalidate() {
    DateTimePicker_Invalidate(d.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (d *TDateTimePicker) PaintTo(DC HDC, X int32, Y int32) {
    DateTimePicker_PaintTo(d.instance, DC , X , Y)
}

// 移除一个控件。
//
// Remove a control.
func (d *TDateTimePicker) RemoveControl(AControl IControl) {
    DateTimePicker_RemoveControl(d.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (d *TDateTimePicker) Realign() {
    DateTimePicker_Realign(d.instance)
}

// 重绘。
//
// Repaint.
func (d *TDateTimePicker) Repaint() {
    DateTimePicker_Repaint(d.instance)
}

// 按比例缩放。
//
// Scale by.
func (d *TDateTimePicker) ScaleBy(M int32, D int32) {
    DateTimePicker_ScaleBy(d.instance, M , D)
}

// 滚动至指定位置。
//
// Scroll by.
func (d *TDateTimePicker) ScrollBy(DeltaX int32, DeltaY int32) {
    DateTimePicker_ScrollBy(d.instance, DeltaX , DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (d *TDateTimePicker) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    DateTimePicker_SetBounds(d.instance, ALeft , ATop , AWidth , AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (d *TDateTimePicker) SetFocus() {
    DateTimePicker_SetFocus(d.instance)
}

// 控件更新。
//
// Update.
func (d *TDateTimePicker) Update() {
    DateTimePicker_Update(d.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (d *TDateTimePicker) BringToFront() {
    DateTimePicker_BringToFront(d.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (d *TDateTimePicker) ClientToScreen(Point TPoint) TPoint {
    return DateTimePicker_ClientToScreen(d.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (d *TDateTimePicker) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return DateTimePicker_ClientToParent(d.instance, Point , CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (d *TDateTimePicker) Dragging() bool {
    return DateTimePicker_Dragging(d.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (d *TDateTimePicker) HasParent() bool {
    return DateTimePicker_HasParent(d.instance)
}

// 隐藏控件。
//
// Hidden control.
func (d *TDateTimePicker) Hide() {
    DateTimePicker_Hide(d.instance)
}

// 发送一个消息。
//
// Send a message.
func (d *TDateTimePicker) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return DateTimePicker_Perform(d.instance, Msg , WParam , LParam)
}

// 刷新控件。
//
// Refresh control.
func (d *TDateTimePicker) Refresh() {
    DateTimePicker_Refresh(d.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (d *TDateTimePicker) ScreenToClient(Point TPoint) TPoint {
    return DateTimePicker_ScreenToClient(d.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (d *TDateTimePicker) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return DateTimePicker_ParentToClient(d.instance, Point , CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (d *TDateTimePicker) SendToBack() {
    DateTimePicker_SendToBack(d.instance)
}

// 显示控件。
//
// Show control.
func (d *TDateTimePicker) Show() {
    DateTimePicker_Show(d.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (d *TDateTimePicker) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return DateTimePicker_GetTextBuf(d.instance, Buffer , BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (d *TDateTimePicker) GetTextLen() int32 {
    return DateTimePicker_GetTextLen(d.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (d *TDateTimePicker) SetTextBuf(Buffer string) {
    DateTimePicker_SetTextBuf(d.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (d *TDateTimePicker) FindComponent(AName string) *TComponent {
    return AsComponent(DateTimePicker_FindComponent(d.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (d *TDateTimePicker) GetNamePath() string {
    return DateTimePicker_GetNamePath(d.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (d *TDateTimePicker) Assign(Source IObject) {
    DateTimePicker_Assign(d.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (d *TDateTimePicker) ClassType() TClass {
    return DateTimePicker_ClassType(d.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (d *TDateTimePicker) ClassName() string {
    return DateTimePicker_ClassName(d.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (d *TDateTimePicker) InstanceSize() int32 {
    return DateTimePicker_InstanceSize(d.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (d *TDateTimePicker) InheritsFrom(AClass TClass) bool {
    return DateTimePicker_InheritsFrom(d.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (d *TDateTimePicker) Equals(Obj IObject) bool {
    return DateTimePicker_Equals(d.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (d *TDateTimePicker) GetHashCode() int32 {
    return DateTimePicker_GetHashCode(d.instance)
}

// 文本类信息。
//
// Text information.
func (d *TDateTimePicker) ToString() string {
    return DateTimePicker_ToString(d.instance)
}

func (d *TDateTimePicker) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    DateTimePicker_AnchorToNeighbour(d.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (d *TDateTimePicker) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    DateTimePicker_AnchorParallel(d.instance, ASide , ASpace , CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (d *TDateTimePicker) AnchorHorizontalCenterTo(ASibling IControl) {
    DateTimePicker_AnchorHorizontalCenterTo(d.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (d *TDateTimePicker) AnchorVerticalCenterTo(ASibling IControl) {
    DateTimePicker_AnchorVerticalCenterTo(d.instance, CheckPtr(ASibling))
}

func (d *TDateTimePicker) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    DateTimePicker_AnchorSame(d.instance, ASide , CheckPtr(ASibling))
}

func (d *TDateTimePicker) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    DateTimePicker_AnchorAsAlign(d.instance, ATheAlign , ASpace)
}

func (d *TDateTimePicker) AnchorClient(ASpace int32) {
    DateTimePicker_AnchorClient(d.instance, ASpace)
}

func (d *TDateTimePicker) ScaleDesignToForm(ASize int32) int32 {
    return DateTimePicker_ScaleDesignToForm(d.instance, ASize)
}

func (d *TDateTimePicker) ScaleFormToDesign(ASize int32) int32 {
    return DateTimePicker_ScaleFormToDesign(d.instance, ASize)
}

func (d *TDateTimePicker) Scale96ToForm(ASize int32) int32 {
    return DateTimePicker_Scale96ToForm(d.instance, ASize)
}

func (d *TDateTimePicker) ScaleFormTo96(ASize int32) int32 {
    return DateTimePicker_ScaleFormTo96(d.instance, ASize)
}

func (d *TDateTimePicker) Scale96ToFont(ASize int32) int32 {
    return DateTimePicker_Scale96ToFont(d.instance, ASize)
}

func (d *TDateTimePicker) ScaleFontTo96(ASize int32) int32 {
    return DateTimePicker_ScaleFontTo96(d.instance, ASize)
}

func (d *TDateTimePicker) ScaleScreenToFont(ASize int32) int32 {
    return DateTimePicker_ScaleScreenToFont(d.instance, ASize)
}

func (d *TDateTimePicker) ScaleFontToScreen(ASize int32) int32 {
    return DateTimePicker_ScaleFontToScreen(d.instance, ASize)
}

func (d *TDateTimePicker) Scale96ToScreen(ASize int32) int32 {
    return DateTimePicker_Scale96ToScreen(d.instance, ASize)
}

func (d *TDateTimePicker) ScaleScreenTo96(ASize int32) int32 {
    return DateTimePicker_ScaleScreenTo96(d.instance, ASize)
}

func (d *TDateTimePicker) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
    DateTimePicker_AutoAdjustLayout(d.instance, AMode , AFromPPI , AToPPI , AOldFormWidth , ANewFormWidth)
}

func (d *TDateTimePicker) FixDesignFontsPPI(ADesignTimePPI int32) {
    DateTimePicker_FixDesignFontsPPI(d.instance, ADesignTimePPI)
}

func (d *TDateTimePicker) ScaleFontsPPI(AToPPI int32, AProportion float64) {
    DateTimePicker_ScaleFontsPPI(d.instance, AToPPI , AProportion)
}

func (d *TDateTimePicker) ArrowShape() TArrowShape {
    return DateTimePicker_GetArrowShape(d.instance)
}

func (d *TDateTimePicker) SetArrowShape(value TArrowShape) {
    DateTimePicker_SetArrowShape(d.instance, value)
}

func (d *TDateTimePicker) AutoAdvance() bool {
    return DateTimePicker_GetAutoAdvance(d.instance)
}

func (d *TDateTimePicker) SetAutoAdvance(value bool) {
    DateTimePicker_SetAutoAdvance(d.instance, value)
}

func (d *TDateTimePicker) AutoButtonSize() bool {
    return DateTimePicker_GetAutoButtonSize(d.instance)
}

func (d *TDateTimePicker) SetAutoButtonSize(value bool) {
    DateTimePicker_SetAutoButtonSize(d.instance, value)
}

func (d *TDateTimePicker) Cascade() bool {
    return DateTimePicker_GetCascade(d.instance)
}

func (d *TDateTimePicker) SetCascade(value bool) {
    DateTimePicker_SetCascade(d.instance, value)
}

func (d *TDateTimePicker) CenturyFrom() uint16 {
    return DateTimePicker_GetCenturyFrom(d.instance)
}

func (d *TDateTimePicker) SetCenturyFrom(value uint16) {
    DateTimePicker_SetCenturyFrom(d.instance, value)
}

func (d *TDateTimePicker) DateDisplayOrder() TDateDisplayOrder {
    return DateTimePicker_GetDateDisplayOrder(d.instance)
}

func (d *TDateTimePicker) SetDateDisplayOrder(value TDateDisplayOrder) {
    DateTimePicker_SetDateDisplayOrder(d.instance, value)
}

func (d *TDateTimePicker) DateSeparator() string {
    return DateTimePicker_GetDateSeparator(d.instance)
}

func (d *TDateTimePicker) SetDateSeparator(value string) {
    DateTimePicker_SetDateSeparator(d.instance, value)
}

func (d *TDateTimePicker) LeadingZeros() bool {
    return DateTimePicker_GetLeadingZeros(d.instance)
}

func (d *TDateTimePicker) SetLeadingZeros(value bool) {
    DateTimePicker_SetLeadingZeros(d.instance, value)
}

func (d *TDateTimePicker) MonthNames() string {
    return DateTimePicker_GetMonthNames(d.instance)
}

func (d *TDateTimePicker) SetMonthNames(value string) {
    DateTimePicker_SetMonthNames(d.instance, value)
}

func (d *TDateTimePicker) ShowMonthNames() bool {
    return DateTimePicker_GetShowMonthNames(d.instance)
}

func (d *TDateTimePicker) SetShowMonthNames(value bool) {
    DateTimePicker_SetShowMonthNames(d.instance, value)
}

func (d *TDateTimePicker) NullInputAllowed() bool {
    return DateTimePicker_GetNullInputAllowed(d.instance)
}

func (d *TDateTimePicker) SetNullInputAllowed(value bool) {
    DateTimePicker_SetNullInputAllowed(d.instance, value)
}

func (d *TDateTimePicker) Options() TDateTimePickerOptions {
    return DateTimePicker_GetOptions(d.instance)
}

func (d *TDateTimePicker) SetOptions(value TDateTimePickerOptions) {
    DateTimePicker_SetOptions(d.instance, value)
}

func (d *TDateTimePicker) ShowCheckBox() bool {
    return DateTimePicker_GetShowCheckBox(d.instance)
}

func (d *TDateTimePicker) SetShowCheckBox(value bool) {
    DateTimePicker_SetShowCheckBox(d.instance, value)
}

// 获取只读。
func (d *TDateTimePicker) ReadOnly() bool {
    return DateTimePicker_GetReadOnly(d.instance)
}

// 设置只读。
func (d *TDateTimePicker) SetReadOnly(value bool) {
    DateTimePicker_SetReadOnly(d.instance, value)
}

func (d *TDateTimePicker) TextForNullDate() string {
    return DateTimePicker_GetTextForNullDate(d.instance)
}

func (d *TDateTimePicker) SetTextForNullDate(value string) {
    DateTimePicker_SetTextForNullDate(d.instance, value)
}

func (d *TDateTimePicker) TimeDisplay() TTimeDisplay {
    return DateTimePicker_GetTimeDisplay(d.instance)
}

func (d *TDateTimePicker) SetTimeDisplay(value TTimeDisplay) {
    DateTimePicker_SetTimeDisplay(d.instance, value)
}

func (d *TDateTimePicker) TimeSeparator() string {
    return DateTimePicker_GetTimeSeparator(d.instance)
}

func (d *TDateTimePicker) SetTimeSeparator(value string) {
    DateTimePicker_SetTimeSeparator(d.instance, value)
}

func (d *TDateTimePicker) TrailingSeparator() bool {
    return DateTimePicker_GetTrailingSeparator(d.instance)
}

func (d *TDateTimePicker) SetTrailingSeparator(value bool) {
    DateTimePicker_SetTrailingSeparator(d.instance, value)
}

func (d *TDateTimePicker) UseDefaultSeparators() bool {
    return DateTimePicker_GetUseDefaultSeparators(d.instance)
}

func (d *TDateTimePicker) SetUseDefaultSeparators(value bool) {
    DateTimePicker_SetUseDefaultSeparators(d.instance, value)
}

func (d *TDateTimePicker) DroppedDown() bool {
    return DateTimePicker_GetDroppedDown(d.instance)
}

func (d *TDateTimePicker) DateTime() time.Time {
    return DateTimePicker_GetDateTime(d.instance)
}

func (d *TDateTimePicker) SetDateTime(value time.Time) {
    DateTimePicker_SetDateTime(d.instance, value)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (d *TDateTimePicker) Align() TAlign {
    return DateTimePicker_GetAlign(d.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (d *TDateTimePicker) SetAlign(value TAlign) {
    DateTimePicker_SetAlign(d.instance, value)
}

// 获取四个角位置的锚点。
func (d *TDateTimePicker) Anchors() TAnchors {
    return DateTimePicker_GetAnchors(d.instance)
}

// 设置四个角位置的锚点。
func (d *TDateTimePicker) SetAnchors(value TAnchors) {
    DateTimePicker_SetAnchors(d.instance, value)
}

func (d *TDateTimePicker) BiDiMode() TBiDiMode {
    return DateTimePicker_GetBiDiMode(d.instance)
}

func (d *TDateTimePicker) SetBiDiMode(value TBiDiMode) {
    DateTimePicker_SetBiDiMode(d.instance, value)
}

func (d *TDateTimePicker) CalAlignment() TDTCalAlignment {
    return DateTimePicker_GetCalAlignment(d.instance)
}

func (d *TDateTimePicker) SetCalAlignment(value TDTCalAlignment) {
    DateTimePicker_SetCalAlignment(d.instance, value)
}

// 获取约束控件大小。
func (d *TDateTimePicker) Constraints() *TSizeConstraints {
    return AsSizeConstraints(DateTimePicker_GetConstraints(d.instance))
}

// 设置约束控件大小。
func (d *TDateTimePicker) SetConstraints(value *TSizeConstraints) {
    DateTimePicker_SetConstraints(d.instance, CheckPtr(value))
}

func (d *TDateTimePicker) Date() time.Time {
    return DateTimePicker_GetDate(d.instance)
}

func (d *TDateTimePicker) SetDate(value time.Time) {
    DateTimePicker_SetDate(d.instance, value)
}

func (d *TDateTimePicker) Time() time.Time {
    return DateTimePicker_GetTime(d.instance)
}

func (d *TDateTimePicker) SetTime(value time.Time) {
    DateTimePicker_SetTime(d.instance, value)
}

// 获取是否选中。
func (d *TDateTimePicker) Checked() bool {
    return DateTimePicker_GetChecked(d.instance)
}

// 设置是否选中。
func (d *TDateTimePicker) SetChecked(value bool) {
    DateTimePicker_SetChecked(d.instance, value)
}

// 获取颜色。
//
// Get color.
func (d *TDateTimePicker) Color() TColor {
    return DateTimePicker_GetColor(d.instance)
}

// 设置颜色。
//
// Set color.
func (d *TDateTimePicker) SetColor(value TColor) {
    DateTimePicker_SetColor(d.instance, value)
}

func (d *TDateTimePicker) DateMode() TDTDateMode {
    return DateTimePicker_GetDateMode(d.instance)
}

func (d *TDateTimePicker) SetDateMode(value TDTDateMode) {
    DateTimePicker_SetDateMode(d.instance, value)
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (d *TDateTimePicker) DoubleBuffered() bool {
    return DateTimePicker_GetDoubleBuffered(d.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (d *TDateTimePicker) SetDoubleBuffered(value bool) {
    DateTimePicker_SetDoubleBuffered(d.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (d *TDateTimePicker) Enabled() bool {
    return DateTimePicker_GetEnabled(d.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (d *TDateTimePicker) SetEnabled(value bool) {
    DateTimePicker_SetEnabled(d.instance, value)
}

// 获取字体。
//
// Get Font.
func (d *TDateTimePicker) Font() *TFont {
    return AsFont(DateTimePicker_GetFont(d.instance))
}

// 设置字体。
//
// Set Font.
func (d *TDateTimePicker) SetFont(value *TFont) {
    DateTimePicker_SetFont(d.instance, CheckPtr(value))
}

func (d *TDateTimePicker) MaxDate() time.Time {
    return DateTimePicker_GetMaxDate(d.instance)
}

func (d *TDateTimePicker) SetMaxDate(value time.Time) {
    DateTimePicker_SetMaxDate(d.instance, value)
}

func (d *TDateTimePicker) MinDate() time.Time {
    return DateTimePicker_GetMinDate(d.instance)
}

func (d *TDateTimePicker) SetMinDate(value time.Time) {
    DateTimePicker_SetMinDate(d.instance, value)
}

// 获取使用父容器颜色。
//
// Get parent color.
func (d *TDateTimePicker) ParentColor() bool {
    return DateTimePicker_GetParentColor(d.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (d *TDateTimePicker) SetParentColor(value bool) {
    DateTimePicker_SetParentColor(d.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (d *TDateTimePicker) ParentDoubleBuffered() bool {
    return DateTimePicker_GetParentDoubleBuffered(d.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (d *TDateTimePicker) SetParentDoubleBuffered(value bool) {
    DateTimePicker_SetParentDoubleBuffered(d.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (d *TDateTimePicker) ParentFont() bool {
    return DateTimePicker_GetParentFont(d.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (d *TDateTimePicker) SetParentFont(value bool) {
    DateTimePicker_SetParentFont(d.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (d *TDateTimePicker) ParentShowHint() bool {
    return DateTimePicker_GetParentShowHint(d.instance)
}

// 设置以父容器的ShowHint属性为准。
func (d *TDateTimePicker) SetParentShowHint(value bool) {
    DateTimePicker_SetParentShowHint(d.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (d *TDateTimePicker) PopupMenu() *TPopupMenu {
    return AsPopupMenu(DateTimePicker_GetPopupMenu(d.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (d *TDateTimePicker) SetPopupMenu(value IComponent) {
    DateTimePicker_SetPopupMenu(d.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (d *TDateTimePicker) ShowHint() bool {
    return DateTimePicker_GetShowHint(d.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (d *TDateTimePicker) SetShowHint(value bool) {
    DateTimePicker_SetShowHint(d.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (d *TDateTimePicker) TabOrder() TTabOrder {
    return DateTimePicker_GetTabOrder(d.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (d *TDateTimePicker) SetTabOrder(value TTabOrder) {
    DateTimePicker_SetTabOrder(d.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (d *TDateTimePicker) TabStop() bool {
    return DateTimePicker_GetTabStop(d.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (d *TDateTimePicker) SetTabStop(value bool) {
    DateTimePicker_SetTabStop(d.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (d *TDateTimePicker) Visible() bool {
    return DateTimePicker_GetVisible(d.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (d *TDateTimePicker) SetVisible(value bool) {
    DateTimePicker_SetVisible(d.instance, value)
}

// 设置控件单击事件。
//
// Set control click event.
func (d *TDateTimePicker) SetOnClick(fn TNotifyEvent) {
    DateTimePicker_SetOnClick(d.instance, fn)
}

// 设置改变事件。
//
// Set changed event.
func (d *TDateTimePicker) SetOnChange(fn TNotifyEvent) {
    DateTimePicker_SetOnChange(d.instance, fn)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (d *TDateTimePicker) SetOnContextPopup(fn TContextPopupEvent) {
    DateTimePicker_SetOnContextPopup(d.instance, fn)
}

func (d *TDateTimePicker) SetOnDropDown(fn TNotifyEvent) {
    DateTimePicker_SetOnDropDown(d.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (d *TDateTimePicker) SetOnEnter(fn TNotifyEvent) {
    DateTimePicker_SetOnEnter(d.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (d *TDateTimePicker) SetOnExit(fn TNotifyEvent) {
    DateTimePicker_SetOnExit(d.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (d *TDateTimePicker) SetOnKeyDown(fn TKeyEvent) {
    DateTimePicker_SetOnKeyDown(d.instance, fn)
}

// 设置键键下事件。
func (d *TDateTimePicker) SetOnKeyPress(fn TKeyPressEvent) {
    DateTimePicker_SetOnKeyPress(d.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (d *TDateTimePicker) SetOnKeyUp(fn TKeyEvent) {
    DateTimePicker_SetOnKeyUp(d.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (d *TDateTimePicker) SetOnMouseEnter(fn TNotifyEvent) {
    DateTimePicker_SetOnMouseEnter(d.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (d *TDateTimePicker) SetOnMouseLeave(fn TNotifyEvent) {
    DateTimePicker_SetOnMouseLeave(d.instance, fn)
}

// 获取依靠客户端总数。
func (d *TDateTimePicker) DockClientCount() int32 {
    return DateTimePicker_GetDockClientCount(d.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (d *TDateTimePicker) DockSite() bool {
    return DateTimePicker_GetDockSite(d.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (d *TDateTimePicker) SetDockSite(value bool) {
    DateTimePicker_SetDockSite(d.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (d *TDateTimePicker) MouseInClient() bool {
    return DateTimePicker_GetMouseInClient(d.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (d *TDateTimePicker) VisibleDockClientCount() int32 {
    return DateTimePicker_GetVisibleDockClientCount(d.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (d *TDateTimePicker) Brush() *TBrush {
    return AsBrush(DateTimePicker_GetBrush(d.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (d *TDateTimePicker) ControlCount() int32 {
    return DateTimePicker_GetControlCount(d.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (d *TDateTimePicker) Handle() HWND {
    return DateTimePicker_GetHandle(d.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (d *TDateTimePicker) ParentWindow() HWND {
    return DateTimePicker_GetParentWindow(d.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (d *TDateTimePicker) SetParentWindow(value HWND) {
    DateTimePicker_SetParentWindow(d.instance, value)
}

func (d *TDateTimePicker) Showing() bool {
    return DateTimePicker_GetShowing(d.instance)
}

// 获取使用停靠管理。
func (d *TDateTimePicker) UseDockManager() bool {
    return DateTimePicker_GetUseDockManager(d.instance)
}

// 设置使用停靠管理。
func (d *TDateTimePicker) SetUseDockManager(value bool) {
    DateTimePicker_SetUseDockManager(d.instance, value)
}

func (d *TDateTimePicker) Action() *TAction {
    return AsAction(DateTimePicker_GetAction(d.instance))
}

func (d *TDateTimePicker) SetAction(value IComponent) {
    DateTimePicker_SetAction(d.instance, CheckPtr(value))
}

func (d *TDateTimePicker) BoundsRect() TRect {
    return DateTimePicker_GetBoundsRect(d.instance)
}

func (d *TDateTimePicker) SetBoundsRect(value TRect) {
    DateTimePicker_SetBoundsRect(d.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (d *TDateTimePicker) ClientHeight() int32 {
    return DateTimePicker_GetClientHeight(d.instance)
}

// 设置客户区高度。
//
// Set client height.
func (d *TDateTimePicker) SetClientHeight(value int32) {
    DateTimePicker_SetClientHeight(d.instance, value)
}

func (d *TDateTimePicker) ClientOrigin() TPoint {
    return DateTimePicker_GetClientOrigin(d.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (d *TDateTimePicker) ClientRect() TRect {
    return DateTimePicker_GetClientRect(d.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (d *TDateTimePicker) ClientWidth() int32 {
    return DateTimePicker_GetClientWidth(d.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (d *TDateTimePicker) SetClientWidth(value int32) {
    DateTimePicker_SetClientWidth(d.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (d *TDateTimePicker) ControlState() TControlState {
    return DateTimePicker_GetControlState(d.instance)
}

// 设置控件状态。
//
// Set control state.
func (d *TDateTimePicker) SetControlState(value TControlState) {
    DateTimePicker_SetControlState(d.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (d *TDateTimePicker) ControlStyle() TControlStyle {
    return DateTimePicker_GetControlStyle(d.instance)
}

// 设置控件样式。
//
// Set control style.
func (d *TDateTimePicker) SetControlStyle(value TControlStyle) {
    DateTimePicker_SetControlStyle(d.instance, value)
}

func (d *TDateTimePicker) Floating() bool {
    return DateTimePicker_GetFloating(d.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (d *TDateTimePicker) Parent() *TWinControl {
    return AsWinControl(DateTimePicker_GetParent(d.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (d *TDateTimePicker) SetParent(value IWinControl) {
    DateTimePicker_SetParent(d.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (d *TDateTimePicker) Left() int32 {
    return DateTimePicker_GetLeft(d.instance)
}

// 设置左边位置。
//
// Set Left position.
func (d *TDateTimePicker) SetLeft(value int32) {
    DateTimePicker_SetLeft(d.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (d *TDateTimePicker) Top() int32 {
    return DateTimePicker_GetTop(d.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (d *TDateTimePicker) SetTop(value int32) {
    DateTimePicker_SetTop(d.instance, value)
}

// 获取宽度。
//
// Get width.
func (d *TDateTimePicker) Width() int32 {
    return DateTimePicker_GetWidth(d.instance)
}

// 设置宽度。
//
// Set width.
func (d *TDateTimePicker) SetWidth(value int32) {
    DateTimePicker_SetWidth(d.instance, value)
}

// 获取高度。
//
// Get height.
func (d *TDateTimePicker) Height() int32 {
    return DateTimePicker_GetHeight(d.instance)
}

// 设置高度。
//
// Set height.
func (d *TDateTimePicker) SetHeight(value int32) {
    DateTimePicker_SetHeight(d.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (d *TDateTimePicker) Cursor() TCursor {
    return DateTimePicker_GetCursor(d.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (d *TDateTimePicker) SetCursor(value TCursor) {
    DateTimePicker_SetCursor(d.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (d *TDateTimePicker) Hint() string {
    return DateTimePicker_GetHint(d.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (d *TDateTimePicker) SetHint(value string) {
    DateTimePicker_SetHint(d.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (d *TDateTimePicker) ComponentCount() int32 {
    return DateTimePicker_GetComponentCount(d.instance)
}

// 获取组件索引。
//
// Get component index.
func (d *TDateTimePicker) ComponentIndex() int32 {
    return DateTimePicker_GetComponentIndex(d.instance)
}

// 设置组件索引。
//
// Set component index.
func (d *TDateTimePicker) SetComponentIndex(value int32) {
    DateTimePicker_SetComponentIndex(d.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (d *TDateTimePicker) Owner() *TComponent {
    return AsComponent(DateTimePicker_GetOwner(d.instance))
}

// 获取组件名称。
//
// Get the component name.
func (d *TDateTimePicker) Name() string {
    return DateTimePicker_GetName(d.instance)
}

// 设置组件名称。
//
// Set the component name.
func (d *TDateTimePicker) SetName(value string) {
    DateTimePicker_SetName(d.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (d *TDateTimePicker) Tag() int {
    return DateTimePicker_GetTag(d.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (d *TDateTimePicker) SetTag(value int) {
    DateTimePicker_SetTag(d.instance, value)
}

// 获取左边锚点。
func (d *TDateTimePicker) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(DateTimePicker_GetAnchorSideLeft(d.instance))
}

// 设置左边锚点。
func (d *TDateTimePicker) SetAnchorSideLeft(value *TAnchorSide) {
    DateTimePicker_SetAnchorSideLeft(d.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (d *TDateTimePicker) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(DateTimePicker_GetAnchorSideTop(d.instance))
}

// 设置顶边锚点。
func (d *TDateTimePicker) SetAnchorSideTop(value *TAnchorSide) {
    DateTimePicker_SetAnchorSideTop(d.instance, CheckPtr(value))
}

// 获取右边锚点。
func (d *TDateTimePicker) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(DateTimePicker_GetAnchorSideRight(d.instance))
}

// 设置右边锚点。
func (d *TDateTimePicker) SetAnchorSideRight(value *TAnchorSide) {
    DateTimePicker_SetAnchorSideRight(d.instance, CheckPtr(value))
}

// 获取底边锚点。
func (d *TDateTimePicker) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(DateTimePicker_GetAnchorSideBottom(d.instance))
}

// 设置底边锚点。
func (d *TDateTimePicker) SetAnchorSideBottom(value *TAnchorSide) {
    DateTimePicker_SetAnchorSideBottom(d.instance, CheckPtr(value))
}

func (d *TDateTimePicker) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(DateTimePicker_GetChildSizing(d.instance))
}

func (d *TDateTimePicker) SetChildSizing(value *TControlChildSizing) {
    DateTimePicker_SetChildSizing(d.instance, CheckPtr(value))
}

// 获取边框间距。
func (d *TDateTimePicker) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(DateTimePicker_GetBorderSpacing(d.instance))
}

// 设置边框间距。
func (d *TDateTimePicker) SetBorderSpacing(value *TControlBorderSpacing) {
    DateTimePicker_SetBorderSpacing(d.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (d *TDateTimePicker) DockClients(Index int32) *TControl {
    return AsControl(DateTimePicker_GetDockClients(d.instance, Index))
}

// 获取指定索引子控件。
func (d *TDateTimePicker) Controls(Index int32) *TControl {
    return AsControl(DateTimePicker_GetControls(d.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (d *TDateTimePicker) Components(AIndex int32) *TComponent {
    return AsComponent(DateTimePicker_GetComponents(d.instance, AIndex))
}

// 获取锚侧面。
func (d *TDateTimePicker) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(DateTimePicker_GetAnchorSide(d.instance, AKind))
}

