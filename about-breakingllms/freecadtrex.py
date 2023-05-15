import FreeCAD, Part

doc = FreeCAD.ActiveDocument

# Simple house function from the previous example
def create_simple_house(position):
    base = create_box("Base", 50, 50, 3, FreeCAD.Placement(position, FreeCAD.Rotation()))
    wall_thickness = 2
    wall_height = 20
    left_wall = create_box("LeftWall", wall_thickness, 50, wall_height, FreeCAD.Placement(position.add(FreeCAD.Vector(0, 0, 3)), FreeCAD.Rotation()))
    right_wall = create_box("RightWall", wall_thickness, 50, wall_height, FreeCAD.Placement(position.add(FreeCAD.Vector(48, 0, 3)), FreeCAD.Rotation()))
    front_wall = create_box("FrontWall", 46, wall_thickness, wall_height, FreeCAD.Placement(position.add(FreeCAD.Vector(2, 0, 3)), FreeCAD.Rotation()))
    back_wall = create_box("BackWall", 46, wall_thickness, wall_height, FreeCAD.Placement(position.add(FreeCAD.Vector(2, 48, 3)), FreeCAD.Rotation()))

    roof = Part.makeLoft([Part.makePolygon([position.add(FreeCAD.Vector(2,2,23)), position.add(FreeCAD.Vector(48,2,23)), position.add(FreeCAD.Vector(25,2,33)), position.add(FreeCAD.Vector(2,2,23))]),
                          Part.makePolygon([position.add(FreeCAD.Vector(2,48,23)), position.add(FreeCAD.Vector(48,48,23)), position.add(FreeCAD.Vector(25,48,33)), position.add(FreeCAD.Vector(2,48,23))])], True)
    roof_obj = doc.addObject("Part::Feature", "Roof")
    roof_obj.Shape = roof

# Create a very simplified T-Rex-like shape
def create_simple_trex(position):
    body = create_box("Body", 20, 8, 10, FreeCAD.Placement(position, FreeCAD.Rotation()))
    head = create_box("Head", 12, 6, 8, FreeCAD.Placement(position.add(FreeCAD.Vector(20, 1, 5)), FreeCAD.Rotation()))
    leg1 = create_box("Leg1", 4, 5, 12, FreeCAD.Placement(position.add(FreeCAD.Vector(5, 0, -12)), FreeCAD.Rotation()))
    leg2 = create_box("Leg2", 4, 5, 12, FreeCAD.Placement(position.add(FreeCAD.Vector(15, 0, -12)), FreeCAD.Rotation()))
    arm1 = create_box("Arm1", 2, 2, 6, FreeCAD.Placement(position.add(FreeCAD.Vector(10, 8, 4)), FreeCAD.Rotation()))
    arm2 = create_box("Arm2", 2, 2, 6, FreeCAD.Placement(position.add(FreeCAD.Vector(10, -4, 4)), FreeCAD.Rotation()))

# Create the house
create_simple_house(FreeCAD.Vector(0, 0, 0))

# Create the simplified T-Rex