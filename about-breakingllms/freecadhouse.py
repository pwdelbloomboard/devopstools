import FreeCAD, Part

doc = FreeCAD.ActiveDocument

def create_box(name, length, width, height, placement):
    box = Part.makeBox(length, width, height)
    box_obj = doc.addObject("Part::Feature", name)
    box_obj.Shape = box
    box_obj.Placement = placement
    return box_obj

# Create the base of the house
base = create_box("Base", 50, 50, 3, FreeCAD.Placement(FreeCAD.Vector(0, 0, 0), FreeCAD.Rotation()))

# Create the walls
wall_thickness = 2
wall_height = 20
left_wall = create_box("LeftWall", wall_thickness, 50, wall_height, FreeCAD.Placement(FreeCAD.Vector(0, 0, 3), FreeCAD.Rotation()))
right_wall = create_box("RightWall", wall_thickness, 50, wall_height, FreeCAD.Placement(FreeCAD.Vector(48, 0, 3), FreeCAD.Rotation()))
front_wall = create_box("FrontWall", 46, wall_thickness, wall_height, FreeCAD.Placement(FreeCAD.Vector(2, 0, 3), FreeCAD.Rotation()))
back_wall = create_box("BackWall", 46, wall_thickness, wall_height, FreeCAD.Placement(FreeCAD.Vector(2, 48, 3), FreeCAD.Rotation()))

# Create the roof
roof = Part.makeLoft([Part.makePolygon([FreeCAD.Vector(2,2,23), FreeCAD.Vector(48,2,23), FreeCAD.Vector(25,2,33), FreeCAD.Vector(2,2,23)]),
                      Part.makePolygon([FreeCAD.Vector(2,48,23), FreeCAD.Vector(48,48,23), FreeCAD.Vector(25,48,33), FreeCAD.Vector(2,48,23)])], True)
roof_obj = doc.addObject("Part::Feature", "Roof")
roof_obj.Shape = roof

doc.recompute()
