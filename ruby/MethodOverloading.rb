class Boy

  def print_details(name)
    "Hey My Name is #{name}"
  end

  def print_details(name,age)
    "Hey My Name is #{name} and #{age}"
  end
end

boy = Boy.new
puts boy.print_details("mukul")
puts boy.print_details("mukul",25)