#### Các thành phần Design Pattern Builder

- Product: đối tượng muốn tạo ra
- Builder Interface: Chỉ chứa các phương thức dùng để gán giá trị cho các thuộc tính trong Product
- Concrete Builder: Tạo ra 1 đối tượng cụ thể cho Product (có thể có nhiều Concrete Builder)
- Director: Liên kết Builder Interface với nhiều Concrete Builder
