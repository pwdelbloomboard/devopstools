import FreeCAD, Part

# Create a box with dimensions 10x10x10
box = Part.makeBox(10, 10, 10)

# Add the box to the active document
doc = FreeCAD.ActiveDocument
box_obj = doc.addObject("Part::Feature", "Box")
box_obj.Shape = box
doc.recompute()