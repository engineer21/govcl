
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TMiniWebview struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewMiniWebview(owner IComponent) *TMiniWebview {
    m := new(TMiniWebview)
    m.instance = MiniWebview_Create(CheckPtr(owner))
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsMiniWebview(obj interface{}) *TMiniWebview {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TMiniWebview{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsMiniWebview.
func MiniWebviewFromInst(inst uintptr) *TMiniWebview {
    return AsMiniWebview(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsMiniWebview.
func MiniWebviewFromObj(obj IObject) *TMiniWebview {
    return AsMiniWebview(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsMiniWebview.
func MiniWebviewFromUnsafePointer(ptr unsafe.Pointer) *TMiniWebview {
    return AsMiniWebview(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (m *TMiniWebview) Free() {
    if m.instance != 0 {
        MiniWebview_Free(m.instance)
        m.instance, m.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (m *TMiniWebview) Instance() uintptr {
    return m.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (m *TMiniWebview) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (m *TMiniWebview) IsValid() bool {
    return m.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (m *TMiniWebview) Is() TIs {
    return TIs(m.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (m *TMiniWebview) As() TAs {
//    return TAs(m.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TMiniWebviewClass() TClass {
    return MiniWebview_StaticClassType()
}

func (m *TMiniWebview) Navigate(AURL string) {
    MiniWebview_Navigate(m.instance, AURL)
}

func (m *TMiniWebview) GoBack() {
    MiniWebview_GoBack(m.instance)
}

func (m *TMiniWebview) GoForward() {
    MiniWebview_GoForward(m.instance)
}

func (m *TMiniWebview) GoHome() {
    MiniWebview_GoHome(m.instance)
}

func (m *TMiniWebview) GoSearch() {
    MiniWebview_GoSearch(m.instance)
}

// 刷新控件。
//
// Refresh control.
func (m *TMiniWebview) Refresh() {
    MiniWebview_Refresh(m.instance)
}

func (m *TMiniWebview) Stop() {
    MiniWebview_Stop(m.instance)
}

// 设置组件边界。
//
// Set component boundaries.
func (m *TMiniWebview) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    MiniWebview_SetBounds(m.instance, ALeft , ATop , AWidth , AHeight)
}

func (m *TMiniWebview) ExecuteScript(AScriptText string, AScriptType string) string {
    return MiniWebview_ExecuteScript(m.instance, AScriptText , AScriptType)
}

func (m *TMiniWebview) ExecuteJS(AScriptText string) string {
    return MiniWebview_ExecuteJS(m.instance, AScriptText)
}

func (m *TMiniWebview) LoadHTML(AStr string) {
    MiniWebview_LoadHTML(m.instance, AStr)
}

// 是否可以获得焦点。
func (m *TMiniWebview) CanFocus() bool {
    return MiniWebview_CanFocus(m.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (m *TMiniWebview) ContainsControl(Control IControl) bool {
    return MiniWebview_ContainsControl(m.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (m *TMiniWebview) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(MiniWebview_ControlAtPos(m.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (m *TMiniWebview) DisableAlign() {
    MiniWebview_DisableAlign(m.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (m *TMiniWebview) EnableAlign() {
    MiniWebview_EnableAlign(m.instance)
}

// 查找子控件。
//
// Find sub controls.
func (m *TMiniWebview) FindChildControl(ControlName string) *TControl {
    return AsControl(MiniWebview_FindChildControl(m.instance, ControlName))
}

func (m *TMiniWebview) FlipChildren(AllLevels bool) {
    MiniWebview_FlipChildren(m.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (m *TMiniWebview) Focused() bool {
    return MiniWebview_Focused(m.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (m *TMiniWebview) HandleAllocated() bool {
    return MiniWebview_HandleAllocated(m.instance)
}

// 插入一个控件。
//
// Insert a control.
func (m *TMiniWebview) InsertControl(AControl IControl) {
    MiniWebview_InsertControl(m.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (m *TMiniWebview) Invalidate() {
    MiniWebview_Invalidate(m.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (m *TMiniWebview) PaintTo(DC HDC, X int32, Y int32) {
    MiniWebview_PaintTo(m.instance, DC , X , Y)
}

// 移除一个控件。
//
// Remove a control.
func (m *TMiniWebview) RemoveControl(AControl IControl) {
    MiniWebview_RemoveControl(m.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (m *TMiniWebview) Realign() {
    MiniWebview_Realign(m.instance)
}

// 重绘。
//
// Repaint.
func (m *TMiniWebview) Repaint() {
    MiniWebview_Repaint(m.instance)
}

// 按比例缩放。
//
// Scale by.
func (m *TMiniWebview) ScaleBy(M int32, D int32) {
    MiniWebview_ScaleBy(m.instance, M , D)
}

// 滚动至指定位置。
//
// Scroll by.
func (m *TMiniWebview) ScrollBy(DeltaX int32, DeltaY int32) {
    MiniWebview_ScrollBy(m.instance, DeltaX , DeltaY)
}

// 设置控件焦点。
//
// Set control focus.
func (m *TMiniWebview) SetFocus() {
    MiniWebview_SetFocus(m.instance)
}

// 控件更新。
//
// Update.
func (m *TMiniWebview) Update() {
    MiniWebview_Update(m.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (m *TMiniWebview) BringToFront() {
    MiniWebview_BringToFront(m.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (m *TMiniWebview) ClientToScreen(Point TPoint) TPoint {
    return MiniWebview_ClientToScreen(m.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (m *TMiniWebview) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return MiniWebview_ClientToParent(m.instance, Point , CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (m *TMiniWebview) Dragging() bool {
    return MiniWebview_Dragging(m.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (m *TMiniWebview) HasParent() bool {
    return MiniWebview_HasParent(m.instance)
}

// 隐藏控件。
//
// Hidden control.
func (m *TMiniWebview) Hide() {
    MiniWebview_Hide(m.instance)
}

// 发送一个消息。
//
// Send a message.
func (m *TMiniWebview) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return MiniWebview_Perform(m.instance, Msg , WParam , LParam)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (m *TMiniWebview) ScreenToClient(Point TPoint) TPoint {
    return MiniWebview_ScreenToClient(m.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (m *TMiniWebview) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return MiniWebview_ParentToClient(m.instance, Point , CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (m *TMiniWebview) SendToBack() {
    MiniWebview_SendToBack(m.instance)
}

// 显示控件。
//
// Show control.
func (m *TMiniWebview) Show() {
    MiniWebview_Show(m.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (m *TMiniWebview) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return MiniWebview_GetTextBuf(m.instance, Buffer , BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (m *TMiniWebview) GetTextLen() int32 {
    return MiniWebview_GetTextLen(m.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (m *TMiniWebview) SetTextBuf(Buffer string) {
    MiniWebview_SetTextBuf(m.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (m *TMiniWebview) FindComponent(AName string) *TComponent {
    return AsComponent(MiniWebview_FindComponent(m.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (m *TMiniWebview) GetNamePath() string {
    return MiniWebview_GetNamePath(m.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (m *TMiniWebview) Assign(Source IObject) {
    MiniWebview_Assign(m.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (m *TMiniWebview) ClassType() TClass {
    return MiniWebview_ClassType(m.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (m *TMiniWebview) ClassName() string {
    return MiniWebview_ClassName(m.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (m *TMiniWebview) InstanceSize() int32 {
    return MiniWebview_InstanceSize(m.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (m *TMiniWebview) InheritsFrom(AClass TClass) bool {
    return MiniWebview_InheritsFrom(m.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (m *TMiniWebview) Equals(Obj IObject) bool {
    return MiniWebview_Equals(m.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (m *TMiniWebview) GetHashCode() int32 {
    return MiniWebview_GetHashCode(m.instance)
}

// 文本类信息。
//
// Text information.
func (m *TMiniWebview) ToString() string {
    return MiniWebview_ToString(m.instance)
}

func (m *TMiniWebview) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    MiniWebview_AnchorToNeighbour(m.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (m *TMiniWebview) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    MiniWebview_AnchorParallel(m.instance, ASide , ASpace , CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (m *TMiniWebview) AnchorHorizontalCenterTo(ASibling IControl) {
    MiniWebview_AnchorHorizontalCenterTo(m.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (m *TMiniWebview) AnchorVerticalCenterTo(ASibling IControl) {
    MiniWebview_AnchorVerticalCenterTo(m.instance, CheckPtr(ASibling))
}

func (m *TMiniWebview) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    MiniWebview_AnchorSame(m.instance, ASide , CheckPtr(ASibling))
}

func (m *TMiniWebview) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    MiniWebview_AnchorAsAlign(m.instance, ATheAlign , ASpace)
}

func (m *TMiniWebview) AnchorClient(ASpace int32) {
    MiniWebview_AnchorClient(m.instance, ASpace)
}

func (m *TMiniWebview) ScaleDesignToForm(ASize int32) int32 {
    return MiniWebview_ScaleDesignToForm(m.instance, ASize)
}

func (m *TMiniWebview) ScaleFormToDesign(ASize int32) int32 {
    return MiniWebview_ScaleFormToDesign(m.instance, ASize)
}

func (m *TMiniWebview) Scale96ToForm(ASize int32) int32 {
    return MiniWebview_Scale96ToForm(m.instance, ASize)
}

func (m *TMiniWebview) ScaleFormTo96(ASize int32) int32 {
    return MiniWebview_ScaleFormTo96(m.instance, ASize)
}

func (m *TMiniWebview) Scale96ToFont(ASize int32) int32 {
    return MiniWebview_Scale96ToFont(m.instance, ASize)
}

func (m *TMiniWebview) ScaleFontTo96(ASize int32) int32 {
    return MiniWebview_ScaleFontTo96(m.instance, ASize)
}

func (m *TMiniWebview) ScaleScreenToFont(ASize int32) int32 {
    return MiniWebview_ScaleScreenToFont(m.instance, ASize)
}

func (m *TMiniWebview) ScaleFontToScreen(ASize int32) int32 {
    return MiniWebview_ScaleFontToScreen(m.instance, ASize)
}

func (m *TMiniWebview) Scale96ToScreen(ASize int32) int32 {
    return MiniWebview_Scale96ToScreen(m.instance, ASize)
}

func (m *TMiniWebview) ScaleScreenTo96(ASize int32) int32 {
    return MiniWebview_ScaleScreenTo96(m.instance, ASize)
}

func (m *TMiniWebview) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
    MiniWebview_AutoAdjustLayout(m.instance, AMode , AFromPPI , AToPPI , AOldFormWidth , ANewFormWidth)
}

func (m *TMiniWebview) FixDesignFontsPPI(ADesignTimePPI int32) {
    MiniWebview_FixDesignFontsPPI(m.instance, ADesignTimePPI)
}

func (m *TMiniWebview) ScaleFontsPPI(AToPPI int32, AProportion float64) {
    MiniWebview_ScaleFontsPPI(m.instance, AToPPI , AProportion)
}

func (m *TMiniWebview) ReadyState() TReadyState {
    return MiniWebview_GetReadyState(m.instance)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (m *TMiniWebview) Align() TAlign {
    return MiniWebview_GetAlign(m.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (m *TMiniWebview) SetAlign(value TAlign) {
    MiniWebview_SetAlign(m.instance, value)
}

// 获取四个角位置的锚点。
func (m *TMiniWebview) Anchors() TAnchors {
    return MiniWebview_GetAnchors(m.instance)
}

// 设置四个角位置的锚点。
func (m *TMiniWebview) SetAnchors(value TAnchors) {
    MiniWebview_SetAnchors(m.instance, value)
}

// 获取约束控件大小。
func (m *TMiniWebview) Constraints() *TSizeConstraints {
    return AsSizeConstraints(MiniWebview_GetConstraints(m.instance))
}

// 设置约束控件大小。
func (m *TMiniWebview) SetConstraints(value *TSizeConstraints) {
    MiniWebview_SetConstraints(m.instance, CheckPtr(value))
}

// 获取控件启用。
//
// Get the control enabled.
func (m *TMiniWebview) Enabled() bool {
    return MiniWebview_GetEnabled(m.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (m *TMiniWebview) SetEnabled(value bool) {
    MiniWebview_SetEnabled(m.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (m *TMiniWebview) Visible() bool {
    return MiniWebview_GetVisible(m.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (m *TMiniWebview) SetVisible(value bool) {
    MiniWebview_SetVisible(m.instance, value)
}

func (m *TMiniWebview) SetOnTitleChange(fn TWebTitleChangeEvent) {
    MiniWebview_SetOnTitleChange(m.instance, fn)
}

func (m *TMiniWebview) SetOnJSExternal(fn TWebJSExternalEvent) {
    MiniWebview_SetOnJSExternal(m.instance, fn)
}

// 获取依靠客户端总数。
func (m *TMiniWebview) DockClientCount() int32 {
    return MiniWebview_GetDockClientCount(m.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (m *TMiniWebview) DockSite() bool {
    return MiniWebview_GetDockSite(m.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (m *TMiniWebview) SetDockSite(value bool) {
    MiniWebview_SetDockSite(m.instance, value)
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (m *TMiniWebview) DoubleBuffered() bool {
    return MiniWebview_GetDoubleBuffered(m.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (m *TMiniWebview) SetDoubleBuffered(value bool) {
    MiniWebview_SetDoubleBuffered(m.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (m *TMiniWebview) MouseInClient() bool {
    return MiniWebview_GetMouseInClient(m.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (m *TMiniWebview) VisibleDockClientCount() int32 {
    return MiniWebview_GetVisibleDockClientCount(m.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (m *TMiniWebview) Brush() *TBrush {
    return AsBrush(MiniWebview_GetBrush(m.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (m *TMiniWebview) ControlCount() int32 {
    return MiniWebview_GetControlCount(m.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (m *TMiniWebview) Handle() HWND {
    return MiniWebview_GetHandle(m.instance)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (m *TMiniWebview) ParentDoubleBuffered() bool {
    return MiniWebview_GetParentDoubleBuffered(m.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (m *TMiniWebview) SetParentDoubleBuffered(value bool) {
    MiniWebview_SetParentDoubleBuffered(m.instance, value)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (m *TMiniWebview) ParentWindow() HWND {
    return MiniWebview_GetParentWindow(m.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (m *TMiniWebview) SetParentWindow(value HWND) {
    MiniWebview_SetParentWindow(m.instance, value)
}

func (m *TMiniWebview) Showing() bool {
    return MiniWebview_GetShowing(m.instance)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (m *TMiniWebview) TabOrder() TTabOrder {
    return MiniWebview_GetTabOrder(m.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (m *TMiniWebview) SetTabOrder(value TTabOrder) {
    MiniWebview_SetTabOrder(m.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (m *TMiniWebview) TabStop() bool {
    return MiniWebview_GetTabStop(m.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (m *TMiniWebview) SetTabStop(value bool) {
    MiniWebview_SetTabStop(m.instance, value)
}

// 获取使用停靠管理。
func (m *TMiniWebview) UseDockManager() bool {
    return MiniWebview_GetUseDockManager(m.instance)
}

// 设置使用停靠管理。
func (m *TMiniWebview) SetUseDockManager(value bool) {
    MiniWebview_SetUseDockManager(m.instance, value)
}

func (m *TMiniWebview) Action() *TAction {
    return AsAction(MiniWebview_GetAction(m.instance))
}

func (m *TMiniWebview) SetAction(value IComponent) {
    MiniWebview_SetAction(m.instance, CheckPtr(value))
}

func (m *TMiniWebview) BiDiMode() TBiDiMode {
    return MiniWebview_GetBiDiMode(m.instance)
}

func (m *TMiniWebview) SetBiDiMode(value TBiDiMode) {
    MiniWebview_SetBiDiMode(m.instance, value)
}

func (m *TMiniWebview) BoundsRect() TRect {
    return MiniWebview_GetBoundsRect(m.instance)
}

func (m *TMiniWebview) SetBoundsRect(value TRect) {
    MiniWebview_SetBoundsRect(m.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (m *TMiniWebview) ClientHeight() int32 {
    return MiniWebview_GetClientHeight(m.instance)
}

// 设置客户区高度。
//
// Set client height.
func (m *TMiniWebview) SetClientHeight(value int32) {
    MiniWebview_SetClientHeight(m.instance, value)
}

func (m *TMiniWebview) ClientOrigin() TPoint {
    return MiniWebview_GetClientOrigin(m.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (m *TMiniWebview) ClientRect() TRect {
    return MiniWebview_GetClientRect(m.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (m *TMiniWebview) ClientWidth() int32 {
    return MiniWebview_GetClientWidth(m.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (m *TMiniWebview) SetClientWidth(value int32) {
    MiniWebview_SetClientWidth(m.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (m *TMiniWebview) ControlState() TControlState {
    return MiniWebview_GetControlState(m.instance)
}

// 设置控件状态。
//
// Set control state.
func (m *TMiniWebview) SetControlState(value TControlState) {
    MiniWebview_SetControlState(m.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (m *TMiniWebview) ControlStyle() TControlStyle {
    return MiniWebview_GetControlStyle(m.instance)
}

// 设置控件样式。
//
// Set control style.
func (m *TMiniWebview) SetControlStyle(value TControlStyle) {
    MiniWebview_SetControlStyle(m.instance, value)
}

func (m *TMiniWebview) Floating() bool {
    return MiniWebview_GetFloating(m.instance)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (m *TMiniWebview) ShowHint() bool {
    return MiniWebview_GetShowHint(m.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (m *TMiniWebview) SetShowHint(value bool) {
    MiniWebview_SetShowHint(m.instance, value)
}

// 获取控件父容器。
//
// Get control parent container.
func (m *TMiniWebview) Parent() *TWinControl {
    return AsWinControl(MiniWebview_GetParent(m.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (m *TMiniWebview) SetParent(value IWinControl) {
    MiniWebview_SetParent(m.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (m *TMiniWebview) Left() int32 {
    return MiniWebview_GetLeft(m.instance)
}

// 设置左边位置。
//
// Set Left position.
func (m *TMiniWebview) SetLeft(value int32) {
    MiniWebview_SetLeft(m.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (m *TMiniWebview) Top() int32 {
    return MiniWebview_GetTop(m.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (m *TMiniWebview) SetTop(value int32) {
    MiniWebview_SetTop(m.instance, value)
}

// 获取宽度。
//
// Get width.
func (m *TMiniWebview) Width() int32 {
    return MiniWebview_GetWidth(m.instance)
}

// 设置宽度。
//
// Set width.
func (m *TMiniWebview) SetWidth(value int32) {
    MiniWebview_SetWidth(m.instance, value)
}

// 获取高度。
//
// Get height.
func (m *TMiniWebview) Height() int32 {
    return MiniWebview_GetHeight(m.instance)
}

// 设置高度。
//
// Set height.
func (m *TMiniWebview) SetHeight(value int32) {
    MiniWebview_SetHeight(m.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (m *TMiniWebview) Cursor() TCursor {
    return MiniWebview_GetCursor(m.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (m *TMiniWebview) SetCursor(value TCursor) {
    MiniWebview_SetCursor(m.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (m *TMiniWebview) Hint() string {
    return MiniWebview_GetHint(m.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (m *TMiniWebview) SetHint(value string) {
    MiniWebview_SetHint(m.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (m *TMiniWebview) ComponentCount() int32 {
    return MiniWebview_GetComponentCount(m.instance)
}

// 获取组件索引。
//
// Get component index.
func (m *TMiniWebview) ComponentIndex() int32 {
    return MiniWebview_GetComponentIndex(m.instance)
}

// 设置组件索引。
//
// Set component index.
func (m *TMiniWebview) SetComponentIndex(value int32) {
    MiniWebview_SetComponentIndex(m.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (m *TMiniWebview) Owner() *TComponent {
    return AsComponent(MiniWebview_GetOwner(m.instance))
}

// 获取组件名称。
//
// Get the component name.
func (m *TMiniWebview) Name() string {
    return MiniWebview_GetName(m.instance)
}

// 设置组件名称。
//
// Set the component name.
func (m *TMiniWebview) SetName(value string) {
    MiniWebview_SetName(m.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (m *TMiniWebview) Tag() int {
    return MiniWebview_GetTag(m.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (m *TMiniWebview) SetTag(value int) {
    MiniWebview_SetTag(m.instance, value)
}

// 获取左边锚点。
func (m *TMiniWebview) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(MiniWebview_GetAnchorSideLeft(m.instance))
}

// 设置左边锚点。
func (m *TMiniWebview) SetAnchorSideLeft(value *TAnchorSide) {
    MiniWebview_SetAnchorSideLeft(m.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (m *TMiniWebview) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(MiniWebview_GetAnchorSideTop(m.instance))
}

// 设置顶边锚点。
func (m *TMiniWebview) SetAnchorSideTop(value *TAnchorSide) {
    MiniWebview_SetAnchorSideTop(m.instance, CheckPtr(value))
}

// 获取右边锚点。
func (m *TMiniWebview) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(MiniWebview_GetAnchorSideRight(m.instance))
}

// 设置右边锚点。
func (m *TMiniWebview) SetAnchorSideRight(value *TAnchorSide) {
    MiniWebview_SetAnchorSideRight(m.instance, CheckPtr(value))
}

// 获取底边锚点。
func (m *TMiniWebview) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(MiniWebview_GetAnchorSideBottom(m.instance))
}

// 设置底边锚点。
func (m *TMiniWebview) SetAnchorSideBottom(value *TAnchorSide) {
    MiniWebview_SetAnchorSideBottom(m.instance, CheckPtr(value))
}

func (m *TMiniWebview) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(MiniWebview_GetChildSizing(m.instance))
}

func (m *TMiniWebview) SetChildSizing(value *TControlChildSizing) {
    MiniWebview_SetChildSizing(m.instance, CheckPtr(value))
}

// 获取边框间距。
func (m *TMiniWebview) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(MiniWebview_GetBorderSpacing(m.instance))
}

// 设置边框间距。
func (m *TMiniWebview) SetBorderSpacing(value *TControlBorderSpacing) {
    MiniWebview_SetBorderSpacing(m.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (m *TMiniWebview) DockClients(Index int32) *TControl {
    return AsControl(MiniWebview_GetDockClients(m.instance, Index))
}

// 获取指定索引子控件。
func (m *TMiniWebview) Controls(Index int32) *TControl {
    return AsControl(MiniWebview_GetControls(m.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (m *TMiniWebview) Components(AIndex int32) *TComponent {
    return AsComponent(MiniWebview_GetComponents(m.instance, AIndex))
}

// 获取锚侧面。
func (m *TMiniWebview) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(MiniWebview_GetAnchorSide(m.instance, AKind))
}

